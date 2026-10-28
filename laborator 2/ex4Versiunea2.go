// in loc sa am mesajul trimis de canal ca status al goroutinei folosesc close pentru a semnala inchiderea trimiterii in canal = procesarea se sfarseste

package main

import (
	"fmt"
	"time"
)

func goroutine(ch chan string) {
	ch <- "mesaj"
	time.Sleep(300 * time.Millisecond)
	close(ch)
}

func main() {
	var ch = make(chan string)

	go goroutine(ch)

	mesaj, ok := <-ch
	if ok {
		fmt.Printf("mesajul a fost trimis: %s", mesaj)
	} else {
		fmt.Println("nu a fost trimis nimic prin canal")
	}
}

// imi place mai mult prima varianta de rezolvare pentru ca se poate vedea concurenta goroutinelor mult mai clar in rulare (dar probabil varianta 2 e rezolvarea "ceruta" ca sa folosim close())

// in prima rezolvare goroutina main are acces la canal si atunci cand goroutina secundara se afla in sleep, de aia in prima versiune de ex 4
// o sa se afiseze cel mai probabil mesajul "goroutina inca se afla in executie" macar o data
