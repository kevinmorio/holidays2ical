package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
)

/// Calculation of the easter date

func easterOffset(year int) int {
	x := year
	k := x / 100
	m := 15 + (3*k)/4 - (8*k+13)/25
	s := 2 - (3*k)/4
	a := x % 19
	d := (19*a + m) % 30
	r := (d + 1/11) / 29
	og := 21 + d - r
	sz := 7 - (x+x/4+s)%7
	oe := 7 - (og-sz)%7
	os := og + oe

	return os
}

func easterDate(year int) time.Time {
	reference := time.Date(year, time.March, 0, 0, 0, 0, 0, time.UTC)
	return reference.AddDate(0, 0, easterOffset(year))
}

/// Holidays with fixed time

type Holiday struct {
	Name map[string]string
	Date time.Time
}

func (h *Holiday) holidayToEvent(lang string) *ics.VEvent {
	event := ics.NewEvent(strings.ToUpper(uuid.NewString()))
	// event.SetAllDayStartAt(h.Date)
	// event.SetAllDayEndAt(h.Date.AddDate(0, 0, 1))

	event.SetProperty(ics.ComponentPropertyDtStart, h.Date.UTC().Format("20060102"), ics.WithValue(string(ics.ValueDataTypeDate)))
	event.SetProperty(ics.ComponentPropertyDtEnd, h.Date.AddDate(0, 0, 1).UTC().Format("20060102"), ics.WithValue(string(ics.ValueDataTypeDate)))

	// event.SetProperty(ics.ComponentPropertyDtStart, h.Date.UTC().Format("20060102"))
	// event.SetProperty(ics.ComponentPropertyDtEnd, h.Date.AddDate(0, 0, 1).UTC().Format("20060102"))

	event.SetTimeTransparency(ics.TransparencyTransparent)
	event.SetSummary(h.Name[lang])
	event.SetDtStampTime(time.Now())

	return event
}

func NewYear(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Neujahr",
			"en-US": "New Year",
		},
		Date: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
	}
}

func Epiphany(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Heilige Drei Könige",
		},
		Date: time.Date(year, 1, 6, 0, 0, 0, 0, time.UTC),
	}
}

func WomensDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Frauentag",
		},
		Date: time.Date(year, 3, 8, 0, 0, 0, 0, time.UTC),
	}
}

func WorkersDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Tag der Arbeit",
		},
		Date: time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC),
	}
}

func AugsburgerHohesFriedensfest(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Augsburger Hohes Friedensfest",
		},
		Date: time.Date(year, 8, 8, 0, 0, 0, 0, time.UTC),
	}
}

func AssumptionOfMary(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Mariä Himmelfahrt",
		},
		Date: time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC),
	}
}

func ChildrensDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Weltkindertag",
		},
		Date: time.Date(year, 9, 20, 0, 0, 0, 0, time.UTC),
	}
}

func GermanUnityDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Tag der Deutschen Einheit",
		},
		Date: time.Date(year, 10, 3, 0, 0, 0, 0, time.UTC),
	}
}

func ReformationDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Reformationstag",
		},
		Date: time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC),
	}
}

func AllSaintsDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Allerheiligen",
		},
		Date: time.Date(year, 11, 1, 0, 0, 0, 0, time.UTC),
	}
}

func ChristmasEve(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Heiligabend",
		},
		Date: time.Date(year, 12, 24, 0, 0, 0, 0, time.UTC),
	}
}

func FristChristmasDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "1. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC),
	}
}

func SecondChristmasDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "2. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC),
	}
}

// Chaning holidays

func MaundyThursday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Gründonnerstag",
		},
		Date: easterDate(year).AddDate(0, 0, -3),
	}
}

func GoodFriday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Karfreitag",
		},
		Date: easterDate(year).AddDate(0, 0, -2),
	}
}

func EasterSunday(year int) Holiday {
	reference := time.Date(year, time.March, 0, 0, 0, 0, 0, time.UTC)
	date := reference.AddDate(0, 0, easterOffset(year))

	return Holiday{
		Name: map[string]string{
			"de-DE": "Ostersonntag",
		},
		Date: date,
	}
}

func EasterMonday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Ostermontag",
		},
		Date: easterDate(year).AddDate(0, 0, 1),
	}
}

func FeastOfTheAscension(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Christi-Himmelfahrt",
		},
		Date: easterDate(year).AddDate(0, 0, 39),
	}
}

func PentecostSunday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Pfingstsonntag",
		},
		Date: easterDate(year).AddDate(0, 0, 49),
	}
}

func PentecostMonday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Pfingstmontag",
		},
		Date: easterDate(year).AddDate(0, 0, 50),
	}
}

func FeastOfCorpusChristi(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Fronleichnam",
		},
		Date: easterDate(year).AddDate(0, 0, 60),
	}
}

func BußUndBettag(year int) Holiday {
	date := time.Date(year, 11, 22, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Wednesday {
			return Holiday{
				Name: map[string]string{
					"de-DE": "Buß- und Bettag",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

var allHolidays = [](func(int) Holiday){
	NewYear,
	Epiphany,
	WomensDay,
	WorkersDay,
	AugsburgerHohesFriedensfest,
	AssumptionOfMary,
	ChildrensDay,
	GermanUnityDay,
	ReformationDay,
	AllSaintsDay,
	ChristmasEve,
	FristChristmasDay,
	SecondChristmasDay,
	MaundyThursday,
	GoodFriday,
	EasterSunday,
	EasterMonday,
	FeastOfTheAscension,
	PentecostSunday,
	PentecostMonday,
	FeastOfCorpusChristi,
	BußUndBettag,
}

func HolidaysForYear(year int) []Holiday {
	holidays := []Holiday{}

	for _, holiday := range allHolidays {
		holidays = append(holidays, holiday(year))
	}

	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Date.Before(holidays[j].Date)
	})

	return holidays
}

var calendarName = map[string]string{
	"de-DE": "Feiertage",
	"en-US": "Holidays",
}

func main() {
	fromYear := flag.Int("from", time.Now().Year(), "year to start from")
	tillYear := flag.Int("till", time.Now().Year(), "year to end")
	lang := flag.String("lang", "de-DE", "the language used for the holidays")

	defaultOutfilePath := fmt.Sprintf("./%s.ics", calendarName[*lang])
	outfilePath := flag.String("outfile", defaultOutfilePath, "the outfile of the calendar")

	flag.Parse()

	cal := ics.NewCalendarFor("-//Kevin Morio//holidays2ics")
	cal.SetCalscale("GREGORIAN")
	cal.SetXWRCalName(calendarName[*lang])

	for year := *fromYear; year <= *tillYear; year++ {
		for _, holiday := range HolidaysForYear(year) {
			cal.AddVEvent(holiday.holidayToEvent(*lang))
		}
	}

	outfile, err := os.Create(*outfilePath)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	if err := cal.SerializeTo(outfile); err != nil {
		panic(err)
	}
	fmt.Printf("Saved calendar to %s\n", *outfilePath)
}
