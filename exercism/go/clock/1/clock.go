// Package clock implements a simple clock that can add and subtract minutes.
package clock

import "fmt"

// Clock defines the hour and the minute fields for our clock.
type Clock struct {
	hour   int
	minute int
}

// New returns a Clock with the given time.
func New(h, m int) Clock {
	c := Clock{}

	c.hour = (h + m/60) % 24
	if c.hour < 0 {
		c.hour += 24
	}

	c.minute = m % 60
	if c.minute < 0 {
		c.minute += 60
		c.hour -= 1
	}

	return c
}

// Add returns the clock with the given number of minutes added to the time.
func (c Clock) Add(m int) Clock {
	c.hour = (c.hour + ((c.minute + m) / 60)) % 24
	c.minute = (c.minute + m) % 60

	return c
}

// Subtract returns the clock with the given number of minutes subtracted from
// the time.
func (c Clock) Subtract(m int) Clock {
	c.hour = (c.hour - (m / 60)) % 24
	if c.hour < 0 {
		c.hour += 24
	}
	c.minute = c.minute - (m % 60)
	if c.minute < 0 {
		c.minute += 60
		c.hour -= 1
		if c.hour < 0 {
			c.hour += 24
		}
	}

	return c
}

// String returns the clock's time as a string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
