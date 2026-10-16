package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func conversie(text string) int {
	nr, er := strconv.Atoi(text)
	if er != nil {
		fmt.Println("Eroare la conversie")
		return -1
	} else {
		return nr
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Err() != nil {
		fmt.Println("Exista o eraore la citire")
	}

	scanner.Scan()
	n := conversie(scanner.Text())
	intSlice := make([]int, 0, n)

	scanner.Scan()
	index := conversie(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		nrConvertit := conversie(scanner.Text())
		intSlice = append(intSlice, nrConvertit)
	}

	primaJumatate := intSlice[0:index]
	douaJumatate := intSlice[index+1 : len(intSlice)]

	intSlice = append(primaJumatate, douaJumatate...)

	fmt.Println(intSlice)
}
