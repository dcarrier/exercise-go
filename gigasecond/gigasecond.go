// Package gigasecond will return your 1G birthday.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4 

// AddGigasecond function will return the given date, time object (t), plus 1G.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(1e18)
	return t
}	

