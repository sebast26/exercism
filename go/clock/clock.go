package clock

import "fmt"

// Clock represents time.
type Clock struct {
	minutes int
}

const dayMinutes = 1440
const hourMinutes = 60

// New creates clock from hours and minutes
func New(h, m int) Clock {
	return Clock{minutes: toMinutes(h, m)}
}

// Add minutes to clock
func (c Clock) Add(m int) Clock {
	c.minutes = toMinutes(c.getHours(), c.getMinutes()+m)
	return c
}

// Subtract minutes from clock
func (c Clock) Subtract(m int) Clock {
	c.minutes = toMinutes(c.getHours(), c.getMinutes()-m)
	return c
}

func toMinutes(h, m int) int {
	mins := (h*hourMinutes + m) % dayMinutes
	if mins < 0 {
		mins = dayMinutes + (mins % dayMinutes)
	}
	return mins
}

func (c Clock) getHours() int {
	return c.minutes / hourMinutes % dayMinutes
}

func (c Clock) getMinutes() int {
	return c.minutes % hourMinutes
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.getHours(), c.getMinutes())
}
