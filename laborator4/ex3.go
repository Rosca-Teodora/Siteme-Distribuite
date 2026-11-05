package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Pachet2 struct {
	x      float64
	maxVot float64 // imp pt conditia de oprire (cea mentionata in lab i.e. diferenta neglijabila dintre min si max)
	minVot float64
}

type Nod3 struct {
	votInitial   float64 // v_i
	votCurent    float64 // notatia din curs pt x(0) = vot initial; x(i) = iteratia i
	votMax       float64
	votMin       float64
	status       bool
	rounds       int
	diametru     int
	epsilon      float64 // cand devine min - max neglijabil!
	canalIntrare []chan Pachet2
	canalIesire  []chan Pachet2
}

func (nod *Nod3) send() { // send(vot)
	for i := 0; i < len(nod.canalIesire); i++ {
		nod.canalIesire[i] <- Pachet2{x: nod.votCurent, maxVot: nod.votMax, minVot: nod.votMin}
	}
}

func (nod *Nod3) receive() []Pachet2 {
	pachete := make([]Pachet2, 0)

	for i := 0; i < len(nod.canalIntrare); i++ {
		canalC := nod.canalIntrare[i]
		pachete = append(pachete, <-canalC)
	}

	return pachete
}

func (nod *Nod3) startIteratieNoua() {
	nod.votMax = nod.votCurent
	nod.votMin = nod.votCurent
}

func (nod *Nod3) transformare(_wg *sync.WaitGroup) {
	defer _wg.Done()
	nod.votCurent = nod.votInitial

	nod.startIteratieNoua()
	nod.send()

	for true {
		nod.rounds++
		received := nod.receive()

		suma := nod.votCurent
		for i := 0; i < len(received); i++ {
			pachetPrimit := received[i]

			suma += pachetPrimit.x

			if pachetPrimit.maxVot > nod.votMax {
				nod.votMax = pachetPrimit.maxVot
			}

			if pachetPrimit.minVot < nod.votMin {
				nod.votMin = pachetPrimit.minVot
			}
		}
		nod.votCurent = (float64(suma)) / (float64(len(received) + 1))

		if nod.rounds%nod.diametru == 0 {
			if nod.votMax-nod.votMin < nod.epsilon {
				if nod.votCurent > 0.5 {
					nod.status = true
				} else {
					nod.status = false
				}
				fmt.Printf("Se voteaza %t!!!!\n", nod.status)

				return
			}
			nod.startIteratieNoua()
		}

		nod.send()
	}
}

func main() {
	_wg := sync.WaitGroup{}
	diametru := 4
	epsilon := 1e-6

	ids := []int{9, 2, 6, 14, 4, 3, 17, 8, 22, 19}
	muchii := [][2]int{
		// exterior
		{9, 2}, {2, 6}, {6, 14}, {14, 4}, {4, 3},
		{3, 17}, {17, 8}, {8, 22}, {22, 19}, {19, 9},
		// interior
		{9, 14}, {2, 4}, {6, 3}, {14, 17}, {4, 8},
		{3, 22}, {17, 19}, {8, 9}, {22, 2}, {19, 6},
	}

	noduri := map[int]*Nod3{}
	for _, id := range ids {
		noduri[id] = &Nod3{votInitial: rand.Float64(), diametru: diametru, epsilon: epsilon}
	}

	for _, m := range muchii {
		ch := make(chan Pachet2, len(ids)+1)
		noduri[m[0]].canalIesire = append(noduri[m[0]].canalIesire, ch)
		noduri[m[1]].canalIntrare = append(noduri[m[1]].canalIntrare, ch)
	}

	_wg.Add(len(ids))
	for _, id := range ids {
		go noduri[id].transformare(&_wg)
	}
	_wg.Wait()

	suma := 0.0
	for _, id := range ids {
		suma += noduri[id].votInitial
	}
	fmt.Printf("\nMedia reala a voturilor initiale = %.6f\n\n", suma/float64(len(ids)))

	for _, id := range ids {
		n := noduri[id]
		decizie := "nu"
		if n.status {
			decizie = "da"
		}
		fmt.Printf("Nodul %d: v = %.3f -> x(%d) = %.6f, max-min = %.1e -> %s\n",
			id, n.votInitial, n.rounds, n.votCurent, n.votMax-n.votMin, decizie)
	}
}
