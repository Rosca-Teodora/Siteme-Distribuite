// NU m-as fi prins de capul meu ca structura asta simuleaza un set daca nu auzeam intrebarea lui Fernando la finalul laboratorului
// util: https://codesignal.com/learn/courses/mastering-complex-data-structures-in-go/lessons/simulating-sets-in-go-using-maps

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	//"sync"
	"os"
)

// func (nod *Nod) transformare(sendCh chan Pachet, recvCh chan Pachet, _wg *sync.WaitGroup, nrCrash int, s int) {
// 	defer _wg.Done()
// 	nod.sendId = nod.nodId

// 	// momentan nu opresc crashul nodurilor...
// 	sansaCrash := rand.intN(100)
// 	if sansaCrash < nrCrash {
// 		return // nodul nu mai e functional deci ies din goroutina
// 	}

// 	for true {
// 		fmt.Printf("Nodul %d trimite: {%d, %t}\n", nod.nodId, nod.sendId, nod.status)
// 		sendCh <- Pachet{nod.sendId, nod.status}

// 		if nod.status {
// 			return
// 		}

// 		pachetPrimit := <-recvCh

// 		if pachetPrimit.done {
// 			sendCh <- pachetPrimit
// 			return
// 		}

// 		if pachetPrimit.id > nod.nodId {
// 			nod.sendId = pachetPrimit.id
// 		} else {
// 			if pachetPrimit.id == nod.nodId {
// 				nod.status = true
// 				nod.sendId = pachetPrimit.id
// 				fmt.Printf("Nodul %d s-a marcat ca lider!!!!\n", nod.nodId)
// 			}
// 			if pachetPrimit.id < nod.nodId {
// 				nod.sendId = -1
// 			}
// 		}
// 	}

// }

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0

	for src := 0; src < n; src++ {
		// BFS din nodul src
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1 // avem un nod la care nu putem ajunge din src
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

// Info reprezinta o pereche (id , valoare ) retinuta in memoria unui nod.
type Info struct {
	ID    int
	Value int
}

// Memory este o multime de token -uri ( reuniunea se face prin inserare ).
type Memory struct {
	elements map[Info]struct{}
}

func NewMemory() *Memory {
	return &Memory{elements: make(map[Info]struct{})}
}

func (m *Memory) Add(inf Info) {
	m.elements[inf] = struct{}{}
}

func (m *Memory) Has(inf Info) bool {
	_, exista := m.elements[inf]
	return exista
}

func (m *Memory) Union(m2 *Memory) *Memory {
	result := NewMemory()
	for v := range m.elements {
		result.Add(v)
	}
	for v := range m2.elements {
		result.Add(v)
	}
	return result
}

type Node struct {
	ID        int
	Neighbors []int
	M         *Memory
}

func main() {
	fisier, err := os.Open("input1.in")
	if err != nil {
		fmt.Println("Eroare la citire")
		return
	}
	defer fisier.Close()

	scanner := bufio.NewScanner(fisier)
	scanner.Scan()
	line := strings.Fields(scanner.Text())
	n, _ := strconv.ParseInt(line[0], 10, 0)
	s, _ := strconv.ParseInt(line[1], 10, 0)

	scanner.Scan()
	line = strings.Fields(scanner.Text())
	nodes := make([]Node, 0, n)
	for i := range line {
		nodId := i
		valoare, _ := strconv.ParseInt(line[i], 10, 0)
		vecini := make([]int, 0)
		mem := NewMemory()
		mem.Add(Info{ID: nodId, Value: int(valoare)})

		nodes = append(nodes, Node{ID: nodId, Neighbors: vecini, M: mem})
	}
}
