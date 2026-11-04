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
	for i := 0; i <= len(nod.canale); i++ {
		conexiuneCurenta := nod.canale[i]
		if conexiuneCurenta.dst == dst {
			conexiuneCurenta.ch <- Pachet{nod.maxId, nod.status}
			return
		}
	}
}

func (nod *Nod2) getVecini() {
	for i := 0; i <= len(nod.canale); i++ {
		conexiuneCurenta := nod.canale[i]
		if conexiuneCurenta.dst == nod.nodId {
			nod.vecini = append(nod.vecini, conexiuneCurenta.dst)
		}
	}
}

func (nod *Nod2) transformare(_wg *sync.WaitGroup) {
	defer _wg.Done()
	nod.maxId = nod.nodId
	nod.getVecini()

	for true {
		nod.rounds++

		for i := 0; i < len(nod.canale); i++ {
			idC := nod.canale[i].returnNodMaxim()

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

	n1 := Nod{nodId: 7, nodDreaptaId: 4, nodStangaId: 1}
	n2 := Nod{nodId: 4, nodDreaptaId: 9, nodStangaId: 7}
	n3 := Nod{nodId: 9, nodDreaptaId: 2, nodStangaId: 4}
	n4 := Nod{nodId: 2, nodDreaptaId: 5, nodStangaId: 9}
	n5 := Nod{nodId: 5, nodDreaptaId: 33, nodStangaId: 2}
	n6 := Nod{nodId: 33, nodDreaptaId: 10, nodStangaId: 5}
	n7 := Nod{nodId: 10, nodDreaptaId: 4, nodStangaId: 33}
	n8 := Nod{nodId: 1, nodDreaptaId: 7, nodStangaId: 10}

	var ch74 = make(chan Pachet, 10)
	var ch49 = make(chan Pachet, 10)
	var ch92 = make(chan Pachet, 10)
	var ch25 = make(chan Pachet, 10)
	var ch533 = make(chan Pachet, 10)
	var ch3310 = make(chan Pachet, 10)
	var ch101 = make(chan Pachet, 10)
	var ch17 = make(chan Pachet, 10)

	_wg.Add(8)
	go n1.transformare(ch74, ch17, &_wg)
	go n2.transformare(ch49, ch74, &_wg)
	go n3.transformare(ch92, ch49, &_wg)
	go n4.transformare(ch25, ch92, &_wg)
	go n5.transformare(ch533, ch25, &_wg)
	go n6.transformare(ch3310, ch533, &_wg)
	go n7.transformare(ch101, ch3310, &_wg)
	go n8.transformare(ch17, ch101, &_wg)
	_wg.Wait()

	// toate canalele unbuffered = blocking
	// sugestie: buffer pt canale (n + 1 pt fiecare canal)
}
