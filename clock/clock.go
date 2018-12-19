package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	hour   int
	minute int
}

func New(hours, minutes int) Clock {
	hour := int(math.Floor(float64(hours*60+minutes)/60)) % 24
	minute := minutes % 60

	if hour < 0 {
		hour += 24
	}

	if minute < 0 {
		minute += 60
	}

	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
