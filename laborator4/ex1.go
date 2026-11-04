package main

import (
	"fmt"
	"sync"
)

type Pachet struct {
	id   int
	done bool
}

type Nod struct {
	nodId, nodDreaptaId, nodStangaId int
	sendId, recvId                   int
	status                           bool
	sendCh, recvCh                   chan Pachet
}

func (nod *Nod) transformare(sendCh chan Pachet, recvCh chan Pachet, _wg *sync.WaitGroup) {
	defer _wg.Done()
	nod.sendId = nod.nodId

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

// func (nod *Nod) addSendCh(sendCh chan Pachet) {
// 	nod.sendCh = sendCh
// }

// func (nod *Nod) addRecvCh(recvCh chan Pachet) {
// 	nod.recvCh = recvCh
// }

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
