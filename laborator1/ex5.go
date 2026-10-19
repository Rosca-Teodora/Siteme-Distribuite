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
	intSlice := make([]string, 0, 50)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := conversie(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		intSlice = append(intSlice, scanner.Text())
	}

	for i := 0; i < len(intSlice)/2; i++ {
		aux := intSlice[i]
		intSlice[i] = intSlice[len(intSlice)-1-i]
		intSlice[len(intSlice)-1-i] = aux
	}

	fmt.Print(intSlice)
}
