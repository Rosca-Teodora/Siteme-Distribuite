package main

import "fmt"

func ex1(ch chan string) {
	ch <- "salut"
}

func main() {
	var ch = make(chan string)

	go ex1(ch)

	var mesaj = <-ch

	fmt.Println(mesaj)
}
