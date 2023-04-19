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

const (
	icalDateFormat string = "20060102"
	ICSFormat      string = "ics"
	StdoutFormat   string = "stdout"
)

func holidayToEvent(h *holidays.Holiday, lang language.Tag) (*ics.VEvent, error) {
	event := ics.NewEvent(strings.ToUpper(uuid.NewString()))

	event.SetProperty(ics.ComponentPropertyDtStart, h.Date.UTC().Format(icalDateFormat), ics.WithValue(string(ics.ValueDataTypeDate)))
	event.SetProperty(ics.ComponentPropertyDtEnd, h.Date.AddDate(0, 0, 1).UTC().Format(icalDateFormat), ics.WithValue(string(ics.ValueDataTypeDate)))

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
	format := flag.String("format", "stdout", "the output format for the holidays (ics|stdout)")
	outfilePath := flag.String("outfile", "Holidays.ics", "the outfile of the calendar")

	flag.Parse()

	langTag, err := language.Parse(*lang)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid language tag '%s'\n", *lang)
	}

	switch *format {
	case ICSFormat:
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
	case StdoutFormat:
		for year := *fromYear; year <= *tillYear; year++ {
			for _, holiday := range holidays.HolidaysForYear(year) {
				fmt.Printf("%s    %s\n", holiday.Date.Format("Mon Jan _2 2006"), holiday.Name[langTag])
			}
		}
	}
}
