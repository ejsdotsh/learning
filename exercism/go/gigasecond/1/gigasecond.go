// Gigasecond expresses time in terms of seconds.
package gigasecond

import (
	"time"
)

// AddGigasecond adds one gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	gigaSecond := 1000000000
	futureDate := t.Add(time.Second * time.Duration(gigaSecond))

	return futureDate
}
