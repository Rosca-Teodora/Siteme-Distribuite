package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// stocheaza deviatia fata de timpul real
type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

// TODO:
// Calculeaza decalajul theta si round -trip -ul delta
func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = ((t2.Sub(t1)) + (t3.Sub(t4))) / 2

	delta = (t4.Sub(t1)) - (t3.Sub(t2))
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {
	// Request
	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.IntN(10))*time.
		Millisecond)
	t2 := s_clock.Now()

	// Procesare intre 5 si 20 ms
	time.Sleep(5*time.Millisecond + time.Duration(rand.IntN(20))*time.Millisecond)
	t3 := s_clock.Now()

	// Response
	time.Sleep(3*time.Millisecond + time.Duration(rand.IntN(10))*time.
		Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	var c1 = Clock{offset: time.Duration(10)}
	var c2 = Clock{offset: time.Duration(20)}

	thetaMin, deltaMin := measure(&c1, &c2)
	i := 1
	for i <= 9 {
		thetaCurent, deltaCurent := measure(&c1, &c2)

		if thetaCurent < thetaMin {
			thetaMin = thetaCurent
		}
		if deltaCurent < deltaMin {
			deltaCurent = deltaMin
		}

		i += 1
	}

	fmt.Println(thetaMin, deltaMin)
}

/*
Rularea celor 2 puncte:
teo@teo:~/Documents/SD/laborator3$ go run ex1PunctulA.go
4.534581ms 15.704643ms
teo@teo:~/Documents/SD/laborator3$ go run ex1PunctulB.go
-4.467439ms 8.852366ms

delta si theta sunt cu mult mai mici dupa mai multe rulari, offset-ul se micsoreaza treptat (cel mai probabil asimptotic catre 0, dar niciodata zero)
*/
