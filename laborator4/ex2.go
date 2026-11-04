package main

import (
	"fmt"
	"sync"
)

type Pachet struct {
	id   int
	done bool
}

type Conexiune struct {
	ch       chan Pachet
	src, dst int // id-urile nodurilor
}

func (conexiune Conexiune) returnNodMaxim() (nodMax int) {
	if conexiune.src >= conexiune.dst {
		return conexiune.src
	}
	return conexiune.dst
}

type Nod2 struct { // am modificat nodul a.i. sa primeasca un slice de canale
	nodId    int
	maxId    int
	status   bool
	rounds   int
	canale   []Conexiune
	vecini   []int // important!!! un nod nu isi cunoaste vecinii de la inceput, ii deduce prin conexiunile pe care le are (ma gandeam ca un sistem nu ar avea de ce sa tina by default vecinii sai in retea) (asta si NU mai stateam sa ii hardcodez pe toti)
	diametru int
}

// si acum chiar am nevoie de metoda pt adaugarea canalului in slice
func (nod *Nod2) addCh(ch Conexiune) {
	nod.canale = append(nod.canale, ch)
}

func (nod *Nod2) send(dst int) { // send(max_id, vecini iesire)
	for i := 0; i < len(nod.canale); i++ {
		conexiuneCurenta := nod.canale[i]
		if conexiuneCurenta.dst == dst {
			conexiuneCurenta.ch <- Pachet{nod.maxId, nod.status}
			return
		}
	}
}

func (nod *Nod2) receive() []int {
	multimeIds := make([]int, 0)

	for i := 0; i < len(nod.canale); i++ {
		conexiuneCurenta := nod.canale[i]
		if conexiuneCurenta.dst == nod.nodId {
			pachet := <-conexiuneCurenta.ch
			multimeIds = append(multimeIds, pachet.id)
		}
	}

	return multimeIds
}

func (nod *Nod2) getVecini() {
	for i := 0; i < len(nod.canale); i++ {
		conexiuneCurenta := nod.canale[i]
		if conexiuneCurenta.src == nod.nodId {
			nod.vecini = append(nod.vecini, conexiuneCurenta.dst)
		}
	}
}

func (nod *Nod2) transformare(_wg *sync.WaitGroup) {
	defer _wg.Done()
	nod.maxId = nod.nodId
	nod.getVecini()

	fmt.Printf("Prima trimitere, nodul %d trimite:\n", nod.nodId)
	for i := range nod.vecini {
		vecinC := nod.vecini[i]
		nod.send(vecinC)
		fmt.Printf("{maxId: %d, vecin: %d}\n", nod.maxId, vecinC)
	}

	for true {
		nod.rounds++
		received := nod.receive()

		for i := 0; i < len(received); i++ {
			idC := received[i]

			if idC > nod.maxId {
				nod.maxId = idC
			}
		}

		if nod.rounds == nod.diametru {
			if nod.maxId == nod.nodId {
				nod.status = true
				fmt.Printf("Nodul %d s-a marcat ca lider!!!!\n", nod.nodId)
			} else {
				nod.status = false
			}
			return
		} else {
			fmt.Printf("Nodul %d incepe sa trimita catre vecini:\n", nod.nodId)
			for i := range nod.vecini {
				vecin := nod.vecini[i]
				nod.send(vecin)
				fmt.Printf("maxId = %d, status = %t\n", nod.maxId, nod.status)
			}
		}
	}
}

func main() {
	_wg := sync.WaitGroup{}

	diam := 4
	ids := []int{9, 2, 6, 14, 4, 3, 17, 8, 22, 19}
	muchii := [][2]int{
		// exterior
		{9, 2}, {2, 6}, {6, 14}, {14, 4}, {4, 3},
		{3, 17}, {17, 8}, {8, 22}, {22, 19}, {19, 9},
		// interior
		{9, 14}, {2, 4}, {6, 3}, {14, 17}, {4, 8},
		{3, 22}, {17, 19}, {8, 9}, {22, 2}, {19, 6},
	}

	noduri := map[int]*Nod2{}
	for _, id := range ids {
		noduri[id] = &Nod2{nodId: id, diametru: diam}
	}

	for _, m := range muchii {
		c := Conexiune{ch: make(chan Pachet, len(ids)+1), src: m[0], dst: m[1]}
		noduri[m[0]].addCh(c) // sursa are canalul ca iesire
		noduri[m[1]].addCh(c) // destinatia il are ca intrare
	}

	_wg.Add(len(ids))
	for _, id := range ids {
		go noduri[id].transformare(&_wg)
	}
	_wg.Wait()
}
