// https://stackoverflow.com/questions/18695346/how-can-i-sort-a-mapstringint-by-its-values

package main

import (
	"fmt"
	"sort"
)

func main() {
	dict := map[string]int{
		"Andrei ": 190,
		"Maria":   165,
		"Ion":     182,
		"Elena":   172,
		"George ": 178,
		"Ana":     165,
	}

	keys := make([]string, 0, len(dict))
	for elem := range dict {
		keys = append(keys, elem)
	}

	sort.Slice(keys, func(i int, j int) bool { return dict[keys[i]] > dict[keys[j]] })

	for _, elem := range keys {
		fmt.Printf("%s, %d\n", elem, dict[elem])
	}
}
