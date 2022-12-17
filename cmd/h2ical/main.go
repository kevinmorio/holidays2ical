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
	"golang.org/x/text/language"
)

var calendarName = holidays.TranslatedString{
	language.German:  "Feiertage",
	language.English: "Holidays",
}

func holidayToEvent(h *holidays.Holiday, lang language.Tag) (*ics.VEvent, error) {
	event := ics.NewEvent(strings.ToUpper(uuid.NewString()))
	// event.SetAllDayStartAt(h.Date)
	// event.SetAllDayEndAt(h.Date.AddDate(0, 0, 1))

	event.SetProperty(ics.ComponentPropertyDtStart, h.Date.UTC().Format("20060102"), ics.WithValue(string(ics.ValueDataTypeDate)))
	event.SetProperty(ics.ComponentPropertyDtEnd, h.Date.AddDate(0, 0, 1).UTC().Format("20060102"), ics.WithValue(string(ics.ValueDataTypeDate)))

	// event.SetProperty(ics.ComponentPropertyDtStart, h.Date.UTC().Format("20060102"))
	// event.SetProperty(ics.ComponentPropertyDtEnd, h.Date.AddDate(0, 0, 1).UTC().Format("20060102"))

	event.SetTimeTransparency(ics.TransparencyTransparent)

	// Consider event name as required
	hName, ok := h.Name[lang]
	if !ok {
		return nil, fmt.Errorf("Name not available for language '%s`", lang)
	}

	// Description is optional
	hDescription := h.Description[lang]

	event.SetSummary(hName)
	event.SetDtStampTime(time.Now())
	event.SetDescription(hDescription)

	return event, nil
}

func main() {
	fromYear := flag.Int("from", time.Now().Year(), "year to start from")
	tillYear := flag.Int("till", time.Now().Year(), "year to end")
	lang := flag.String("lang", "de", "the language used for the holidays")
	outfilePath := flag.String("outfile", "Holidays.ics", "the outfile of the calendar")

	flag.Parse()

	langTag, err := language.Parse(*lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid language tag '%s'\n", *lang)
	}

	cal := ics.NewCalendarFor("-//Kevin Morio//holidays2ics")
	cal.SetCalscale("GREGORIAN")
	cal.SetXWRCalName(calendarName[langTag])

	for year := *fromYear; year <= *tillYear; year++ {
		for _, holiday := range holidays.HolidaysForYear(year) {
			event, err := holidayToEvent(&holiday, langTag)
			if err == nil {
				cal.AddVEvent(event)
			}
		}
	}

	outfile, err := os.Create(*outfilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	defer outfile.Close()

	if err := cal.SerializeTo(outfile); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Printf("Saved calendar to %s\n", *outfilePath)
}
