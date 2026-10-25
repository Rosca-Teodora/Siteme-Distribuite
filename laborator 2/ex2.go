package main

import "fmt"

func ex2(ch chan string) {
	var mesaj = <-ch
	fmt.Println(mesaj)
}

func main() {
	var ch = make(chan string)

	go ex2(ch)

	ch <- "merge?"
}

// initial pusesem adaugarea in canal prima, nu inceperea gorutinei si normal ca mi-a dat deadlock (all goroutines asleep)
