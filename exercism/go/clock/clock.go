// Package clock contains functionality related to creating and manipulating a 24-hour clock
package clock

import "fmt"

// Clock is a 24-hour representation of the time
type Clock struct {
	Hours   int
	Minutes int
}

var (
	minutesPerHour = 60
	minutesPerDay  = 24 * minutesPerHour
)

// New returns an instance of a clock object
func New(hours int, minutes int) Clock {
	minutes = hours*minutesPerHour + minutes
	minutes %= minutesPerDay
	if minutes < 0 {
		minutes += minutesPerDay
	}

	return Clock{minutes / minutesPerHour, minutes % minutesPerHour}
}

// String converts a Clock to type string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

// Add returns the new time with added minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.Hours, c.Minutes+minutes)
}

// Subtract returns the new time with subtracted minutes
func (c Clock) Subtract(minutes int) Clock {
	return New(c.Hours, c.Minutes-minutes)
}
