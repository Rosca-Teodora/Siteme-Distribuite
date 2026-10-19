// pt citire si parsare string: https://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var arr [100]int
	scanner := bufio.NewScanner(os.Stdin)
	freq := make(map[int][]int)

	// scanner.Scan()
	// n := conversie(scanner.Text())

	// scanner.Scan()
	// sum := conversie(scanner.Text())

	// for i := 0; i < n; i++ {
	// 	scanner.Scan()
	// 	arr[i] = conversie(scanner.Text())
	// }

	scanner.Scan()
	line := strings.Fields(scanner.Text())
	n := conversie(line[0])
	sum := conversie(line[1])

	scanner.Scan()
	line2 := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		arr[i] = conversie(line2[i])
	}

	//fmt.Println(arr)

	for index, nrC := range arr {
		_, ok := freq[nrC]
		if !ok {
			freq[nrC] = []int{index}
		} else {
			freq[nrC] = append(freq[nrC], index)
		}
	}

	for nrC, indexiC := range freq {
		diferenta := sum - nrC
		if diferenta == nrC {
			if len(indexiC) > 1 {
				fmt.Printf("%d %d\n", indexiC[0], indexiC[1])
				delete(freq, nrC)
			}
		} else {
			indexiDif, ok := freq[diferenta]
			if ok {
				fmt.Printf("%d %d\n", indexiC[0], indexiDif[0])
				delete(freq, nrC)
			}
		}
	}
}
