package main

// pur si simplu simulez ca ei se intrec si trimit functiei main ca s-au intrecut, fiecare robot e o goroutina

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func robot(ch chan string, idRobot string) {
	time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
	ch <- "A ajuns la final robotul cu id-ul" + idRobot
}

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)

	go robot(ch1, "1")
	go robot(ch2, "2")

	select {
	case castiga := <-ch1:
		fmt.Println(castiga)
	case castiga := <-ch2:
		fmt.Println(castiga)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Cursa a durat prea mult, n-a castigat niciun robot")
	}
}
