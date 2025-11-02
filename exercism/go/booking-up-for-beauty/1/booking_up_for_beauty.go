package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    d, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    t := time.Now()
	if t.Before(d) {
        return false
    }
    return true
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    if t.Hour() >= 12 && t.Hour() < 18 {
    	return true
    }
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    msg := `You have an appointment on %s, %s %d, %d, at %d:%d.`
    a := Schedule(date)
    h, m, _ := a.Clock()
	
    return fmt.Sprintf(msg, a.Weekday().String(), a.Month().String(), a.Day(), a.Year(), h, m)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Now()
    a := time.Date(t.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
    return a
}
