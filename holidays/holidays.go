package holidays

import (
	"sort"
	"time"

	"golang.org/x/text/language"
)

type (
	TranslatedString map[language.Tag]string

	Holiday struct {
		Name        TranslatedString
		Date        time.Time
		Description TranslatedString
	}
)

func NewYear(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Neujahrstag",
			language.English: "New Year",
		},
		Date: time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func Epiphany(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Heilige Drei Könige",
		},
		Date: time.Date(year, 1, 6, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Feiertag in Baden-Württemberg, Bayern, Sachsen-Anhalt0",
		},
	}
}

func ValentinesDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Valentinstag",
		},
		Date: time.Date(year, 2, 14, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func Rosenmontag(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Rosenmontag",
		},
		Date: easterDate(year).AddDate(0, 0, -48),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func ShrowveTuesday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Faschingsdienstag",
			language.English: "Shrove Tuesday",
		},
		Date: easterDate(year).AddDate(0, 0, -47),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func AshWednesday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Aschermittwoch",
			language.English: "Ash Wednesday",
		},
		Date: easterDate(year).AddDate(0, 0, -46),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func WomensDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Internationaler Frauentag",
		},
		Date: time.Date(year, 3, 8, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func StartOfDST(year int) Holiday {
	date := time.Date(year, 3, 31, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: TranslatedString{
					language.German: "Beginn der Sommerzeit",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func PalmSunday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Palmsonntag",
			language.English: "Palm Sunday",
		},
		Date: easterDate(year).AddDate(0, 0, -7),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func MaundyThursday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Gründonnerstag",
		},
		Date: easterDate(year).AddDate(0, 0, -3),
		Description: TranslatedString{
			language.German: "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func GoodFriday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Karfreitag",
		},
		Date: easterDate(year).AddDate(0, 0, -2),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func HolySaturday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Karsamstag",
			language.English: "Holy Saturday",
		},
		Date: easterDate(year).AddDate(0, 0, -1),
		Description: TranslatedString{
			language.German: "Gedenktag in Bayern, Hessen, Niedersachsen, Saarland, Rheinland-Pfalz",
		},
	}
}

func Easter(year int) Holiday {
	reference := time.Date(year, time.March, 0, 0, 0, 0, 0, time.UTC)
	date := reference.AddDate(0, 0, easterOffset(year))

	return Holiday{
		Name: TranslatedString{
			language.German:  "Ostern",
			language.English: "Easter",
		},
		Date: date,
		Description: TranslatedString{
			language.German: "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func EasterMonday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Ostermontag",
		},
		Date: easterDate(year).AddDate(0, 0, 1),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func WorkersDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Tag der Arbeit",
		},
		Date: time.Date(year, 5, 1, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func VictoryInEuropeDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Jahrestag der Befreiung vom Nationalsozialismus",
			language.English: "Victory in Europe Day",
		},
		Date: time.Date(year, 5, 8, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag in Berlin, Brandenburg, Bremen, Mecklenburg-Vorpommern, Thüringen",
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
				Name: TranslatedString{
					language.German:  "Muttertag",
					language.English: "Mother's Day",
				},
				Date: date,
				Description: TranslatedString{
					language.German: "Gedenktag",
				},
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func FeastOfTheAscension(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Christi Himmelfahrt",
		},
		Date: easterDate(year).AddDate(0, 0, 39),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func FathersDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Vatertag",
			language.English: "Father's Day",
		},
		Date: easterDate(year).AddDate(0, 0, 39),
	}
}

func Pentecost(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Pfingsten",
		},
		Date: easterDate(year).AddDate(0, 0, 49),
		Description: TranslatedString{
			language.German: "Gedenktag in Baden-Württemberg, Bayern, Berlin, Brandenburg, Bremen, Hamburg, Hessen, Mecklenburg-Vorpommern, Niedersachsen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Sachsen-Anhalt, Schleswig-Holstein, Thüringen",
		},
	}
}

func PentecostMonday(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Pfingstmontag",
		},
		Date: easterDate(year).AddDate(0, 0, 50),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func FeastOfCorpusChristi(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Fronleichnam",
		},
		Date: easterDate(year).AddDate(0, 0, 60),
		Description: TranslatedString{
			language.German: "Feiertag in Baden-Württemberg, Bayern, Hessen, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland, Sachsen, Thüringen",
		},
	}
}

func AugsburgerHohesFriedensfest(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Augsburger Hohes Friedensfest",
		},
		Date: time.Date(year, 8, 8, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Feiertag in Bayern",
		},
	}
}

func AssumptionOfMary(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Mariä Himmelfahrt",
		},
		Date: time.Date(year, 8, 15, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag in Bayern, Saarland, Sachsen, Thüringen",
		},
	}
}

func ChildrensDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Weltkindertag",
		},
		Date: time.Date(year, 9, 20, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Feiertag in Thüringen",
		},
	}
}

func GermanUnityDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Tag der Deutschen Einheit",
		},
		Date: time.Date(year, 10, 3, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func EndOfDST(year int) Holiday {
	date := time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: TranslatedString{
					language.German: "Ende der Sommerzeit",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func ReformationDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Reformationstag",
		},
		Date: time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Feiertag in Brandenburg, Mecklenburg-Vorpommern, Sachsen, Sachsen-Anhalt, Thüringen, Schleswig-Holstein, Hamburg, Niedersachsen, Bremen",
		},
	}
}

func Halloween(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Halloween",
			language.English: "Halloween",
		},
		Date: time.Date(year, 10, 31, 0, 0, 0, 0, time.UTC),
	}
}

func AllSaintsDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Allerheiligen",
		},
		Date: time.Date(year, 11, 1, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Feiertag in Baden-Württemberg, Bayern, Nordrhein-Westfalen, Rheinland-Pfalz, Saarland",
		},
	}
}

func StMartinsDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "St. Martin",
			language.English: "St. Martin's Day",
		},
		Date: time.Date(year, 11, 11, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func BußUndBettag(year int) Holiday {
	date := time.Date(year, 11, 22, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Wednesday {
			return Holiday{
				Name: TranslatedString{
					language.German: "Buß- und Bettag",
				},
				Date: date,
				Description: TranslatedString{
					language.German: "Feiertag in Sachsen",
				},
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func Volkstrauertag(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Volkstrauertag",
			language.English: "Volkstrauertag",
		},
		Date: Totensonntag(year).Date.AddDate(0, 0, -7),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func Totensonntag(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Totensonntag",
			language.English: "Totensonntag",
		},
		Date: FirstAdvent(year).Date.AddDate(0, 0, -7),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func SaintNicholasDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German:  "Nikolaustag",
			language.English: "Saint Nicholas Day",
		},
		Date: time.Date(year, 12, 6, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gedenktag",
		},
	}
}

func FirstAdvent(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "1. Advent",
		},
		Date: SecondAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func SecondAdvent(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "2. Advent",
		},
		Date: ThirdAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func ThirdAdvent(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "3. Advent",
		},
		Date: FourthAdvent(year).Date.AddDate(0, 0, -7),
	}
}

func FourthAdvent(year int) Holiday {
	date := time.Date(year, 12, 24, 0, 0, 0, 0, time.UTC)

	for {
		if date.Weekday() == time.Sunday {
			return Holiday{
				Name: TranslatedString{
					language.German: "4. Advent",
				},
				Date: date,
			}
		}
		date = date.AddDate(0, 0, -1)
	}
}

func ChristmasEve(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Heiligabend",
		},
		Date: time.Date(year, 12, 24, 0, 0, 0, 0, time.UTC),
	}
}

func FirstChristmasDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "1. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func SecondChristmasDay(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "2. Weihnachtsfeiertag",
		},
		Date: time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC),
		Description: TranslatedString{
			language.German: "Gesetzlicher Feiertag",
		},
	}
}

func Silvester(year int) Holiday {
	return Holiday{
		Name: TranslatedString{
			language.German: "Silvester",
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
