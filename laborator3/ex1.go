package main

import "time"

type Clock struct {
	offset time.Duration
}

func (c *Clock) Now() time.Time {
	return time.Now().Add(c.offset)
}

// TODO: calculeaza decalajul theta si intarzierea delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = ((t2.Sub(t1)) + (t3.Sub(t4))) / 2

	delta = (t4.Sub(t1)) - (t3.Sub(t2))
	return theta, delta
}

func main() {
	i := 1
	for i <= 10 {
		theta, delta = estimate()

		i += 1
	}
}
