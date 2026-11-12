package main

import (
	"fmt"
	"sync"
)

type Pachet struct {
	id   int
	v    int
}

func (p Pachet) equals(b Pachet) {
	if p.id == b.id && p.v == b.v && 
}

type Nod struct {
	nodId, nodDreaptaId, nodStangaId int
	sendId, recvId                   int
	status                           bool
	multimeNod                       []Pachet
	rounds                           int // variabila t = initial 0 
}

func (nod *Nod) transformare(sendCh chan Pachet, recvCh chan Pachet, _wg *sync.WaitGroup, nrCrash int, s int) {
	defer _wg.Done()
	nod.sendId = nod.nodId

	// momentan nu opresc crashul nodurilor...
	sansaCrash := rand.intN(100)
	if sansaCrash < nrCrash {
		return // nodul nu mai e functional deci ies din goroutina 
	}

	for true {
		fmt.Printf("Nodul %d trimite: {%d, %t}\n", nod.nodId, nod.sendId, nod.status)
		sendCh <- Pachet{nod.sendId, nod.status}

		if nod.status {
			return
		}

		pachetPrimit := <-recvCh

		if pachetPrimit.done {
			sendCh <- pachetPrimit
			return
		}

		if pachetPrimit.id > nod.nodId {
			nod.sendId = pachetPrimit.id
		} else {
			if pachetPrimit.id == nod.nodId {
				nod.status = true
				nod.sendId = pachetPrimit.id
				fmt.Printf("Nodul %d s-a marcat ca lider!!!!\n", nod.nodId)
			}
			if pachetPrimit.id < nod.nodId {
				nod.sendId = -1
			}
		}
	}

}

func main() {
	_wg := sync.WaitGroup{}
	s := 3

	n1 := Nod{nodId: 7, nodDreaptaId: 4, nodStangaId: 1, rounds: 0}
	n2 := Nod{nodId: 4, nodDreaptaId: 9, nodStangaId: 7, rounds: 0}
	n3 := Nod{nodId: 9, nodDreaptaId: 2, nodStangaId: 4, rounds: 0}
	n4 := Nod{nodId: 2, nodDreaptaId: 5, nodStangaId: 9, rounds: 0}
	n5 := Nod{nodId: 5, nodDreaptaId: 33, nodStangaId: 2, rounds: 0}
	n6 := Nod{nodId: 33, nodDreaptaId: 10, nodStangaId: 5, rounds: 0}
	n7 := Nod{nodId: 10, nodDreaptaId: 4, nodStangaId: 33, rounds: 0}
	n8 := Nod{nodId: 1, nodDreaptaId: 7, nodStangaId: 10, rounds: 0}

	var ch74 = make(chan Pachet, 10)
	var ch49 = make(chan Pachet, 10)
	var ch92 = make(chan Pachet, 10)
	var ch25 = make(chan Pachet, 10)
	var ch533 = make(chan Pachet, 10)
	var ch3310 = make(chan Pachet, 10)
	var ch101 = make(chan Pachet, 10)
	var ch17 = make(chan Pachet, 10)

	_wg.Add(8)
	go n1.transformare(ch74, ch17, &_wg, 2, s)
	go n2.transformare(ch49, ch74, &_wg, 2, s)
	go n3.transformare(ch92, ch49, &_wg, 2, s)
	go n4.transformare(ch25, ch92, &_wg, 2, s)
	go n5.transformare(ch533, ch25, &_wg, 2, s)
	go n6.transformare(ch3310, ch533, &_wg, 2, s)
	go n7.transformare(ch101, ch3310, &_wg, 2, s)
	go n8.transformare(ch17, ch101, &_wg, 2, s)
	_wg.Wait()

	// toate canalele unbuffered = blocking
	// sugestie: buffer pt canale (n + 1 pt fiecare canal)
}
