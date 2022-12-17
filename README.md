<p align="center">
    <img src="assets/logo_black.svg#gh-light-mode-only" alt="holidays2ical">
    <img src="assets/logo_white.svg#gh-dark-mode-only" alt="holidays2ical"><br>
    <img src="https://img.shields.io/github/license/kevinmorio/holidays2ical" alt="license"></br>
    A command-line tool to generate iCal calendars for public holidays and special days.
</p>

### Installation

``` shell
go install github.com/kevinmorio/holidays2ical/cmd/h2ical@latest
```

### Usage

``` shell
Usage of h2ical:
  -from int
    	year to start from (default 2022)
  -lang string
    	the language used for the holidays (default "de")
  -outfile string
    	the outfile of the calendar (default "Holidays.ics")
  -till int
    	year to end (default 2022)
```