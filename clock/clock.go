// Package clock implements interfaces for dealing with wall-clock time
package clock

import "fmt"

// Clock represents the state of a wall clock
type Clock int

// New creates a new 24-hour Clock given a number of hours and minutes
func New(hours, minutes int) Clock {
	return Clock(clamp24hours(hours*60 + minutes))
}

// String converts a clock into a zero-padded time in the format HH:MM
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours(), c.Minutes())
}

// Add adds a number of minutes to a clock
func (c Clock) Add(minutes int) Clock {
	return New(0, int(c)+minutes)
}

// Subtract removes a number of minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, int(c)-minutes)
}

// Hours returns the number of hours on a clock
func (c Clock) Hours() int {
	return int(c) / 60
}

// Minutes returns the position of the minute hand on a clock
func (c Clock) Minutes() int {
	return int(c) % 60
}

func clamp24hours(minutes int) int {
	total := minutes % 1440

	if total < 0 {
		total += 1440
	}

	return total
}
