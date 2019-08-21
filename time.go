package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		t := time.Now()
		fmt.Println(t)
		fmt.Println(t.Unix())

		fmt.Println(t.Format("20060102150405"))
		fmt.Println(t.Format("2006-01-02 15:04:05"))
		fmt.Println(t.Format("15:04:05"))
		fmt.Println(t.Format("15:04:05.000"))

		// Ymdhis
		ymdhis := fmt.Sprintf("%d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
		fmt.Println(ymdhis)
	*/

	t1, _ := time.Parse("20060102150405", "20130203195400")
	fmt.Println(t1.Format("20060102150405"))

	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value

	// 默认的开始时间 0001-01-01 00:00:00
	startTime := time.Time{}.Format("2006-01-02 15:04:05")
	fmt.Println(startTime)
}
