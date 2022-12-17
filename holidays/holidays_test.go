package holidays

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/text/language"
)

func TestHolidays(t *testing.T) {
	testCases := []struct {
		fn   func(int) Holiday
		year int
		want time.Time
	}{
		{NewYear, 2021, time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
		{Epiphany, 2021, time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC)},
		{ValentinesDay, 2021, time.Date(2021, 2, 14, 0, 0, 0, 0, time.UTC)},
		{Rosenmontag, 2021, time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC)},
		{ShrowveTuesday, 2021, time.Date(2021, 2, 16, 0, 0, 0, 0, time.UTC)},
		{AshWednesday, 2021, time.Date(2021, 2, 17, 0, 0, 0, 0, time.UTC)},
		{WomensDay, 2021, time.Date(2021, 3, 8, 0, 0, 0, 0, time.UTC)},
		{StartOfDST, 2021, time.Date(2021, 3, 28, 0, 0, 0, 0, time.UTC)},
		{PalmSunday, 2021, time.Date(2021, 3, 28, 0, 0, 0, 0, time.UTC)},
		{MaundyThursday, 2021, time.Date(2021, 4, 1, 0, 0, 0, 0, time.UTC)},
		{GoodFriday, 2021, time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)},
		{HolySaturday, 2021, time.Date(2021, 4, 3, 0, 0, 0, 0, time.UTC)},
		{Easter, 2021, time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC)},
		{EasterMonday, 2021, time.Date(2021, 4, 5, 0, 0, 0, 0, time.UTC)},
		{WorkersDay, 2021, time.Date(2021, 5, 1, 0, 0, 0, 0, time.UTC)},
		{VictoryInEuropeDay, 2021, time.Date(2021, 5, 8, 0, 0, 0, 0, time.UTC)},
		{MothersDay, 2021, time.Date(2021, 5, 9, 0, 0, 0, 0, time.UTC)},
		{FeastOfTheAscension, 2021, time.Date(2021, 5, 13, 0, 0, 0, 0, time.UTC)},
		{FathersDay, 2021, time.Date(2021, 5, 13, 0, 0, 0, 0, time.UTC)},
		{Pentecost, 2021, time.Date(2021, 5, 23, 0, 0, 0, 0, time.UTC)},
		{PentecostMonday, 2021, time.Date(2021, 5, 24, 0, 0, 0, 0, time.UTC)},
		{FeastOfCorpusChristi, 2021, time.Date(2021, 6, 3, 0, 0, 0, 0, time.UTC)},
		{AugsburgerHohesFriedensfest, 2021, time.Date(2021, 8, 8, 0, 0, 0, 0, time.UTC)},
		{AssumptionOfMary, 2021, time.Date(2021, 8, 15, 0, 0, 0, 0, time.UTC)},
		{ChildrensDay, 2021, time.Date(2021, 9, 20, 0, 0, 0, 0, time.UTC)},
		{GermanUnityDay, 2021, time.Date(2021, 10, 3, 0, 0, 0, 0, time.UTC)},
		{EndOfDST, 2021, time.Date(2021, 10, 31, 0, 0, 0, 0, time.UTC)},
		{ReformationDay, 2021, time.Date(2021, 10, 31, 0, 0, 0, 0, time.UTC)},
		{Halloween, 2021, time.Date(2021, 10, 31, 0, 0, 0, 0, time.UTC)},
		{AllSaintsDay, 2021, time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC)},
		{StMartinsDay, 2021, time.Date(2021, 11, 11, 0, 0, 0, 0, time.UTC)},
		{BußUndBettag, 2021, time.Date(2021, 11, 17, 0, 0, 0, 0, time.UTC)},
		{Volkstrauertag, 2021, time.Date(2021, 11, 14, 0, 0, 0, 0, time.UTC)},
		{Totensonntag, 2021, time.Date(2021, 11, 21, 0, 0, 0, 0, time.UTC)},
		{SaintNicholasDay, 2021, time.Date(2021, 12, 6, 0, 0, 0, 0, time.UTC)},
		{FirstAdvent, 2021, time.Date(2021, 11, 28, 0, 0, 0, 0, time.UTC)},
		{SecondAdvent, 2021, time.Date(2021, 12, 5, 0, 0, 0, 0, time.UTC)},
		{ThirdAdvent, 2021, time.Date(2021, 12, 12, 0, 0, 0, 0, time.UTC)},
		{FourthAdvent, 2021, time.Date(2021, 12, 19, 0, 0, 0, 0, time.UTC)},
		{ChristmasEve, 2021, time.Date(2021, 12, 24, 0, 0, 0, 0, time.UTC)},
		{FirstChristmasDay, 2021, time.Date(2021, 12, 25, 0, 0, 0, 0, time.UTC)},
		{SecondChristmasDay, 2021, time.Date(2021, 12, 26, 0, 0, 0, 0, time.UTC)},
		{Silvester, 2021, time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC)},
	}

	for _, tc := range testCases {
		got := tc.fn(tc.year)
		t.Run(fmt.Sprintf("%s in %d", got.Name[language.German], tc.year), func(t *testing.T) {
			if got.Date != tc.want {
				t.Errorf("got %s; want %s", got.Date.Format("2006-01-02"), tc.want.Format("2006-01-02"))
			}
		})
	}
}
