// Package clock implements a simple clock that can add and subtract minutes.
package clock

import "fmt"

// Clock is an integer.
type Clock int

// New returns a Clock with the given time.
func New(h, m int) Clock {
	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}

	return c
}

// Add returns a clock with the number of minutes added to the time.
func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

// Subtract returns a clock with the number of minutes subtracted from the time.
func (c Clock) Subtract(m int) Clock {
	return New(0, int(c)-m)
}

// String returns the clock's time as a string.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
