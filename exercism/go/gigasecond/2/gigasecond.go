// Gigasecond expresses time in terms of seconds.
package gigasecond

import (
	"time"
)

// AddGigasecond adds one gigasecond to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
