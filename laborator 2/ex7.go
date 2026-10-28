package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// struct pt puncte 2D si intr-un timp random o data la niste intervale generam niste puncte pe care le trimitem de la un robobot la main
type Punct struct {
	x int
	y int
}

func (p *Punct) adaugaX(x int) {
	p.x += x
}

func (p *Punct) adaugaY(y int) {
	p.y += y
}

// fiecare robot incepe mereu de la zero si poate merge intr-o singura iteratie in ambele directii x, y (deci poate merge diagonal)
// in fiecare directie se misca cu o valoare random pana cand trece de destinatie
func robot(ch chan Punct, idRobot string, destinatie Punct) {
	coordonateCurente := Punct{x: 0, y: 0}

	for coordonateCurente.x < destinatie.x || coordonateCurente.y < destinatie.y {
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)

		if destinatie.x > coordonateCurente.x {
			addX := rand.IntN(destinatie.x / 2)
			coordonateCurente.adaugaX(addX)
		}

		if coordonateCurente.y < destinatie.y {
			addY := rand.IntN(destinatie.y / 2)
			coordonateCurente.adaugaY(addY)
		}

		ch <- coordonateCurente
	}

	close(ch)
}

func main() {
	var ch1 = make(chan Punct)
	var ch2 = make(chan Punct)
	destinatie := Punct{300, 300}

	go robot(ch1, "1", destinatie)
	go robot(ch2, "2", destinatie)

	for true {
		select {
		case coord1, ok1 := <-ch1:
			{
				if !ok1 {
					fmt.Printf("Robotul 1 a castigat cursa\n")
					return
				}
				fmt.Printf("Coordonatele robotului 1 sunt: %d, %d\n", coord1.x, coord1.y)
			}
		case coord2, ok2 := <-ch2:
			{
				if !ok2 {
					fmt.Printf("Robotul 2 a castigat cursa\n")
					return
				}
				fmt.Printf("Coordonatele robotului 2 sunt: %d, %d\n", coord2.x, coord2.y)
			}
		}
	}
}
