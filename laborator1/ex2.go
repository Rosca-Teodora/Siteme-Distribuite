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
	var intArr [100]int

	for i := 0; i <= n; i++ {
		scanner.Scan()
		nrConvertit := conversie(scanner.Text())
		intArr[i] = nrConvertit
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j <= n-1; j++ {
			if intArr[i] > intArr[j] {
				aux := intArr[i]
				intArr[i] = intArr[j]
				intArr[j] = aux
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(intArr[i])
		fmt.Print(" ")
	}
}
