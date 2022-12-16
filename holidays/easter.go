package holidays

import "time"

/// Calculation of the easter date

func easterOffset(year int) int {
	x := year
	k := x / 100
	m := 15 + (3*k)/4 - (8*k+13)/25
	s := 2 - (3*k)/4
	a := x % 19
	d := (19*a + m) % 30
	r := (d + a/11) / 29
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
