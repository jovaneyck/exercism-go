package clock

import "fmt"

type Clock struct {
	minutes int
}

var minutesInADay int = 24 * 60

func New(hours int, minutes int) Clock {
	totalMinutes := (60*hours + minutes) % minutesInADay
	if totalMinutes <= 0 {
		totalMinutes = (minutesInADay + totalMinutes) % minutesInADay
	}

	return Clock{totalMinutes}
}

func (c Clock) Add(minutes int) Clock {
	return New(0, c.minutes+minutes)
}

func (c Clock) toMinutesHours() (int, int) {
	h := c.minutes / 60
	m := c.minutes % 60
	return h, m
}

func (c Clock) String() string {
	m, h := c.toMinutesHours()
	return fmt.Sprintf("%02.f:%02.f", float32(m), float32(h))
}
