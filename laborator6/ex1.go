// NU m-as fi prins de capul meu ca structura asta simuleaza un set daca nu auzeam intrebarea lui Fernando la finalul laboratorului
// util: https://codesignal.com/learn/courses/mastering-complex-data-structures-in-go/lessons/simulating-sets-in-go-using-maps

package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
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

func gambleCrash(s int, sCurent *atomic.Int32) bool {
	if int(sCurent.Load()) >= s {
		return false
	}

	p := rand.IntN(100) //
	if p <= 1 {
		return true
	}
	return false
}

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

func (m *Memory) Sort() []Info {
	mem := make([]Info, 0, len(m.elements))

	for i := range m.elements {
		mem = append(mem, i)
	}
	slices.SortFunc(mem, func(a, b Info) int { return a.ID - b.ID })

	return mem
}

type Node struct {
	ID        int
	Neighbors []int
	M         *Memory
	rounds    int
	crashed   bool
}

func (nod *Node) transformare(recv []chan *Memory, send []chan *Memory, memorieFinala *Memory, s int, sCurent *atomic.Int32, maximRunde int, wg *sync.WaitGroup) {
	defer wg.Done()

	for true {
		for _, ch := range send { // Broadcast(Mi(t))
			ch <- nod.M
		}

		if gambleCrash(s, sCurent) { // nodul a dat crash
			sCurent.Add(1)

			// inchide canalele
			for _, ch := range send {
				close(ch)
			}
			nod.crashed = true

			return
		}

		for _, ch := range recv {
			msj, ok := <-ch
			if ok { // daca nu atunci "panic: runtime error: invalid memory address or nil pointer dereference" fiindca iau si canale inchise
				nod.M = nod.M.Union(msj)
			}
		}

		if !(nod.rounds < maximRunde) {
			return
		}

		nod.rounds += 1
	}
}

func stringToInt(c string) int {
	cInt64, _ := strconv.ParseInt(c, 10, 0)
	return int(cInt64)
}

func startFlooding(n int, adj [][]int, nodes []Node, s int) *Memory {
	var wg sync.WaitGroup
	var sCurent atomic.Int32
	memFin := NewMemory()
	maximRunde := (s + 1) * diametru(adj)

	// initializare matrice de canale
	canale := make([][]chan *Memory, n)
	for i := range n {
		canale[i] = make([]chan *Memory, n)
		for j := range n {
			canale[i][j] = make(chan *Memory, maximRunde+1)
		}
	}

	for i := range n {
		var send []chan *Memory
		var recv []chan *Memory

		for _, j := range adj[i] {
			send = append(send, canale[i][j])
		}

		for k := range n {
			for _, j := range adj[k] {
				if j == i {
					recv = append(recv, canale[k][j])
				}
			}
		}

		wg.Add(1)
		go nodes[i].transformare(recv, send, memFin, s, &sCurent, maximRunde, &wg)
	}

	wg.Wait()

	for i := range nodes {
		if !nodes[i].crashed {
			return nodes[i].M // initial luam o valoare dar nu o dadeam ca pointer deci nu se schimba si asa e mai usor sa iau pur si simplu elementele din memoria primului nod care nu a picat
		}
	}
	return NewMemory() // "missing return" real nu ar tb sa se intample pe date care verifica conditiile dar compiler
}

func main() {
	var adj [][]int

	fisier, err := os.Open("input1.in")
	if err != nil {
		fmt.Println("Eroare la citire")
		return
	}
	defer fisier.Close()

	scanner := bufio.NewScanner(fisier)
	scanner.Scan()
	line := strings.Fields(scanner.Text())

	n := stringToInt(line[0])
	s := stringToInt(line[1])

	scanner.Scan()
	line = strings.Fields(scanner.Text())
	nodes := make([]Node, 0, n)
	for i := range line {
		nodId := i
		valoare, _ := strconv.ParseInt(line[i], 10, 0)
		vecini := make([]int, 0)
		mem := NewMemory()
		mem.Add(Info{ID: nodId, Value: int(valoare)})

		nodes = append(nodes, Node{ID: nodId, Neighbors: vecini, M: mem, rounds: 0, crashed: false})
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		line = strings.Fields(scanner.Text())

		convertedLine := make([]int, 0)
		for j := range line {
			conv := stringToInt(line[j])
			convertedLine = append(convertedLine, conv)
		}
		adj = append(adj, convertedLine)
	}

	//fmt.Println(adj)
	fmt.Println(diametru(adj))

	memFin := startFlooding(n, adj, nodes, s)
	for _, i := range memFin.Sort() {
		fmt.Printf("(%d, %d)", i.ID, i.Value)
	}
	fmt.Print("\n")
	for i := range nodes {
		if nodes[i].crashed {
			fmt.Printf("nodul %d a dat crash\n", nodes[i].ID)
		}
	}

}
