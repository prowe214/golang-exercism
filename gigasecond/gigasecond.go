// Package gigasecond performs gigasecond functions on a time
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond should return a date which is one gigasecond
// after the passed in time t
func AddGigasecond(t time.Time) time.Time {

	// convert t to unix seconds, add one gigasecond,
	// convert back to UTC
	t = t.Add(1000000000 * time.Second)

	return t
}
