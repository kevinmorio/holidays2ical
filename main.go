package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
	"github.com/kevinmorio/holidays2ical/holidays"
)

var calendarName = map[string]string{
	"de-DE": "Feiertage",
	"en-US": "Holidays",
}

func holidayToEvent(h *holidays.Holiday, lang string) *ics.VEvent {
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
	event.SetDescription(h.Description[lang])

	return event
}

func main() {
	fromYear := flag.Int("from", time.Now().Year(), "year to start from")
	tillYear := flag.Int("till", time.Now().Year(), "year to end")
	lang := flag.String("lang", "de-DE", "the language used for the holidays")

	defaultOutfilePath := fmt.Sprintf("%s.ics", calendarName[*lang])
	outfilePath := flag.String("outfile", defaultOutfilePath, "the outfile of the calendar")

	flag.Parse()

	cal := ics.NewCalendarFor("-//Kevin Morio//holidays2ics")
	cal.SetCalscale("GREGORIAN")
	cal.SetXWRCalName(calendarName[*lang])

	for year := *fromYear; year <= *tillYear; year++ {
		for _, holiday := range holidays.HolidaysForYear(year) {
			cal.AddVEvent(holidayToEvent(&holiday, *lang))
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
