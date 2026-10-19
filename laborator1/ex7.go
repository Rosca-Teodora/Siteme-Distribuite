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
	strings := make([]string, 0, 100)
	scanner := bufio.NewScanner(os.Stdin)
	freq := make(map[string]int)

	scanner.Scan()
	n := conversie(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		strings = append(strings, scanner.Text())
	}

	for _, word := range strings {
		_, ok := freq[word]
		if !ok {
			freq[word] = 1
		} else {
			freq[word] += 1
		}
	}

	fmt.Print(freq)
}
