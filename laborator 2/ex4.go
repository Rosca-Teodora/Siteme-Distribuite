package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan bool) {
	ch <- false

	time.Sleep(100 * time.Millisecond)

	ch <- true
}

func main() {
	var ch = make(chan bool)

	go goroutine(ch)

	for true {
		var goroutineStatus = <-ch
		if goroutineStatus {
			fmt.Println("s-a terminat executia goroutinei")
			break
		} else {
			fmt.Println("goroutina inca se afla in executie")
		}
	}
}
