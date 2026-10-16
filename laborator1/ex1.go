package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func estePrim(nr int) bool {
	if nr <= 1 {
		return false
	}
	if nr%2 == 0 {
		return false
	}

	for i := 3; i < nr; i += 2 {
		if nr%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Err() != nil {
		fmt.Println("Exista o eroare la citire")
	}

	scanner.Scan()
	text := scanner.Text()
	nr, er := strconv.Atoi(text)
	if er != nil {
		fmt.Printf("Eroare la conversie")
	}

	if estePrim(nr) {
		fmt.Println("Este prim")
	} else {
		fmt.Println("Nu este prim")
	}

}
