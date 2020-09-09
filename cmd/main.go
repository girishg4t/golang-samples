package main

import (
	"fmt"
	"time"
)

// rangeDate returns a date range function over start date to end date inclusive.
// After the end of the range, the range function returns a zero date,
// date.IsZero() is true.
func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func date() {
	holidayList := []string{
		"2019-Aug-12",
		"2019-Aug-15",
		"2019-Sep-02",
		"2019-Sep-10",
		"2019-Oct-02",
		"2019-Oct-08",
		"2019-Oct-28",
		"2019-Nov-12",
		"2019-Dec-25",
		"2020-Feb-21",
		"2020-Mar-10",
		"2020-Apr-02",
		"2020-Apr-06",
		"2020-Apr-14",
		"2020-May-01",
		"2020-May-25",
	}
	date := time.Now()
	start := date.AddDate(0, -12, 0)
	end := date.AddDate(0, 0, 0)
	fmt.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		if !Contains(holidayList, date.Format("2006-Jan-02")) {
			continue
		} else {
			fmt.Println(date.Format("2006-Jan-02"))
		}

	}
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
