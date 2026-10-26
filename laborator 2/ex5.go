package main

import "fmt"

func send(ch chan string) {
	ch <- "un mesaj"
}

// initial incercasem:
//
// func receive(ch chan string) string {
// 	var mesaj = <-ch
// 	return mesaj
// }
//
// si in main pusesem:
// x := go receive(ch)
//
// dar normal ca primeam eroare (in primul rand sintactica si in al doilea rand pentru ca sunt actiuni contradictorii)
// relevant am gasit pe stackoverflow o explicatie: https://stackoverflow.com/questions/20945069/catching-return-values-from-goroutines

func main() {
	var ch = make(chan string)

	go send(ch)

	var x string
	go func() {
		x = <-ch
	}()

	for true {
		if len(x) > 0 {
			fmt.Println(x)
			break
		} else {
			fmt.Println("inca nu s-a trimis niciun mesaj")
		}
	}
}
