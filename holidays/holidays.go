package holidays

import (
	"sort"
	"time"
)

type Holiday struct {
	Name        map[string]string
	Date        time.Time
	Description map[string]string
}

func NewYear(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Neujahrstag",
			"en-US": "New Year",
		},
		Date: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func Epiphany(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Heilige Drei Könige",
		},
		Date: time.Date(year, 1, 6, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Feiertag in Baden-Württemberg, Bayern, Sachsen-Anhalt0",
		},
	}
}

func ValentinesDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Valentinstag",
		},
		Date: time.Date(year, 2, 14, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func Rosenmontag(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Rosenmontag",
		},
		Date: easterDate(year).AddDate(0, 0, -48),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func ShrowveTuesday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Faschingsdienstag",
			"en-US": "Shrove Tuesday",
		},
		Date: easterDate(year).AddDate(0, 0, -47),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func AshWednesday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Aschermittwoch",
			"en-US": "Ash Wednesday",
		},
		Date: easterDate(year).AddDate(0, 0, -46),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func WomensDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Internationaler Frauentag",
		},
		Date: time.Date(year, 3, 8, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func StartOfDST(year int) Holiday {
	date := time.Date(year, 3, 31, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: map[string]string{
					"de-DE": "Beginn der Sommerzeit",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func PalmSunday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Palmsonntag",
			"en-US": "Palm Sunday",
		},
		Date: easterDate(year).AddDate(0, 0, -7),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func MaundyThursday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Gründonnerstag",
		},
		Date: easterDate(year).AddDate(0, 0, -3),
		Description: map[string]string{
			"de-DE": "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func GoodFriday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Karfreitag",
		},
		Date: easterDate(year).AddDate(0, 0, -2),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func HolySaturday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Karsamstag",
			"en-US": "Holy Saturday",
		},
		Date: easterDate(year).AddDate(0, 0, -1),
		Description: map[string]string{
			"de-DE": "Gedenktag in Bayern, Hessen, Niedersachsen, Saarland, Rheinland-Pfalz",
		},
	}
}

func Easter(year int) Holiday {
	reference := time.Date(year, time.March, 0, 0, 0, 0, 0, time.UTC)
	date := reference.AddDate(0, 0, easterOffset(year))

	return Holiday{
		Name: map[string]string{
			"de-DE": "Ostern",
			"en-US": "Easter",
		},
		Date: date,
		Description: map[string]string{
			"de-DE": "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func EasterMonday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Ostermontag",
		},
		Date: easterDate(year).AddDate(0, 0, 1),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func WorkersDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Tag der Arbeit",
		},
		Date: time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func VictoryInEuropeDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Jahrestag der Befreiung vom Nationalsozialismus",
			"en-US": "Victory in Europe Day",
		},
		Date: time.Date(year, 5, 8, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag in Berlin, Brandenburg, Bremen, Mecklenburg-Vorpommern, Thüringen",
		},
	}
}

// Second Sunday of May
func MothersDay(year int) Holiday {
	var count int
	date := time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			count++
		}
		if count == 2 {
			return Holiday{
				Name: map[string]string{
					"de-DE": "Muttertag",
					"en-US": "Mother's Day",
				},
				Date: date,
				Description: map[string]string{
					"de-DE": "Gedenktag",
				},
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func FeastOfTheAscension(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Christi Himmelfahrt",
		},
		Date: easterDate(year).AddDate(0, 0, 39),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func FathersDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Vatertag",
			"en-EN": "Father's Day",
		},
		Date: easterDate(year).AddDate(0, 0, 39),
	}
}

func Pentecost(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Pfingsten",
		},
		Date: easterDate(year).AddDate(0, 0, 49),
		Description: map[string]string{
			"de-DE": "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func PentecostMonday(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Pfingstmontag",
		},
		Date: easterDate(year).AddDate(0, 0, 50),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func FeastOfCorpusChristi(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Fronleichnam",
		},
		Date: easterDate(year).AddDate(0, 0, 60),
		Description: map[string]string{
			"de-DE": "Feiertag in Baden-Württemberg, Bayern, Hessen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Thüringen",
		},
	}
}

func AugsburgerHohesFriedensfest(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Augsburger Hohes Friedensfest",
		},
		Date: time.Date(year, 8, 8, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Feiertag in Bayern",
		},
	}
}

func AssumptionOfMary(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Mariä Himmelfahrt",
		},
		Date: time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag in Bayern, Saarland, Sachsen, Thüringen",
		},
	}
}

func ChildrensDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Weltkindertag",
		},
		Date: time.Date(year, 9, 20, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Feiertag in Thüringen",
		},
	}
}

func GermanUnityDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Tag der Deutschen Einheit",
		},
		Date: time.Date(year, 10, 3, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func EndOfDST(year int) Holiday {
	date := time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: map[string]string{
					"de-DE": "Ende der Sommerzeit",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func ReformationDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Reformationstag",
		},
		Date: time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Feiertag in Brandenburg, Mecklenburg-Vorpommern, Sachsen, Sachsen-Anhalt, Thüringen, Schleswig-Holstein, Hamburg, Niedersachsen, Bremen",
		},
	}
}

func Halloween(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE":  "Halloween",
			"en-USE": "Halloween",
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
		Description: map[string]string{
			"de-DE": "Feiertag in Baden-Württemberg, Bayern, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland",
		},
	}
}

func StMartinsDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE":  "St. Martin",
			"en-USE": "St. Martin's Day",
		},
		Date: time.Date(year, 11, 11, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
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
				Description: map[string]string{
					"de-DE": "Feiertag in Sachsen",
				},
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func Volkstrauertag(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Volkstrauertag",
			"en-US": "Volkstrauertag",
		},
		Date: Totensonntag(year).Date.AddDate(0, 0, -7),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func Totensonntag(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Totensonntag",
			"en-US": "Totensonntag",
		},
		Date: FirstAdvent(year).Date.AddDate(0, 0, -7),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func SaintNicholasDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Nikolaustag",
			"en-US": "Saint Nicholas Day",
		},
		Date: time.Date(year, 12, 6, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gedenktag",
		},
	}
}

func FirstAdvent(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "1. Advent",
		},
		Date: SecondAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func SecondAdvent(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "2. Advent",
		},
		Date: ThirdAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func ThirdAdvent(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "3. Advent",
		},
		Date: FourthAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func FourthAdvent(year int) Holiday {
	date := time.Date(year, 12, 24, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: map[string]string{
					"de-DE": "4. Advent",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
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

func FirstChristmasDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "1. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func SecondChristmasDay(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "2. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC),
		Description: map[string]string{
			"de-DE": "Gesetzlicher Feiertag",
		},
	}
}

func Silvester(year int) Holiday {
	return Holiday{
		Name: map[string]string{
			"de-DE": "Silvester",
		},
		Date: time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC),
	}
}

var allHolidays = [](func(int) Holiday){
	NewYear,
	Epiphany,
	ValentinesDay,
	Rosenmontag,
	ShrowveTuesday,
	AshWednesday,
	WomensDay,
	StartOfDST,
	PalmSunday,
	MaundyThursday,
	GoodFriday,
	HolySaturday,
	Easter,
	EasterMonday,
	WorkersDay,
	VictoryInEuropeDay,
	MothersDay,
	FeastOfTheAscension,
	FathersDay,
	Pentecost,
	PentecostMonday,
	FeastOfCorpusChristi,
	AugsburgerHohesFriedensfest,
	AssumptionOfMary,
	ChildrensDay,
	GermanUnityDay,
	EndOfDST,
	ReformationDay,
	Halloween,
	AllSaintsDay,
	StMartinsDay,
	BußUndBettag,
	Volkstrauertag,
	Totensonntag,
	SaintNicholasDay,
	FirstAdvent,
	SecondAdvent,
	ThirdAdvent,
	FourthAdvent,
	ChristmasEve,
	FirstChristmasDay,
	SecondChristmasDay,
	Silvester,
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
