// Package clock creates time without dates, prints the time, and handles all negative numbers.
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	total int
}

// New function is the constructor for a new clock.
func New(hour, minute int) Clock {
	total := (60 * hour) + minute
	total = CheckTotal(total)
	return Clock{total}
}

// String function formats and prints the clock object with proper time format (hh:mm).
func (c Clock) String() string {
	hour := c.total / 60
	minute := c.total % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Add function increments the clock by a given amount of minutes and returns to total.
func (c Clock) Add(minutes int) Clock {
	total := c.total + minutes
	total = CheckTotal(total)
	return Clock{total}
}

// CheckTotal function ensures that all numbers over 24 hours carry over and that it handles negative numbers properly.
func CheckTotal(total int) int {
	for total >= 1440 || total < 0 {
		if total >= 1440 {
			total -= 1440
		} else if total < 0 {
			total += 1440
		}

	}
	return total
}

