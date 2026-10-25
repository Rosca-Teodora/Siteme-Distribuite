package main

import "fmt"

func trimiteMesaj(ch chan string) {
	ch <- "mesaj trimis"
}

func main() {
	var ch = make(chan string, 2)

	go trimiteMesaj(ch)
	go trimiteMesaj(ch)

	var m1 = <-ch
	fmt.Println(m1)

	var m2 = <-ch
	fmt.Println(m2)
}
