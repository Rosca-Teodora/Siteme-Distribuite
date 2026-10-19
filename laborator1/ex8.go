package main

import (
	"fmt"
)

func main() {
	angajati := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		" Cristian Vasile": 2,
		"Ioana Enache":     1,
	}

	dep := make(map[int][]string)
	for nume, id_map_initial := range angajati {
		_, ok := dep[id_map_initial]
		if !ok {
			dep[id_map_initial] = make([]string, 0, 5)
			dep[id_map_initial] = append(dep[id_map_initial], nume)
		} else {
			dep[id_map_initial] = append(dep[id_map_initial], nume)
		}
	}

	for dep_id, string_nume := range dep {
		fmt.Printf("Departament %d: %s\n", dep_id, string_nume)
	}
}
