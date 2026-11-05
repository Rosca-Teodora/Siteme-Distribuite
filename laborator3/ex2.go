// 		P1
// 	 / 		\
// 	|		 P2
// 	 \		/
// 		P3

package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"sync"
	"time"
)

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter++
}
func (l *LamportClock) Send() int {
	l.Tick()
	ts := l.counter
	return ts
}
func (l *LamportClock) Receive(ts int) {
	if ts > l.counter {
		l.counter = ts
	}
	l.counter++
}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string // " internal ", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func proces(id int, chInt <-chan Message, chOut chan<- Message, processEvents []Event, _wg *sync.WaitGroup) {
	defer _wg.Done()

	clock := LamportClock{counter: 0}

	k := rand.IntN(5)
	for i := 0; i <= k; i++ {
		time.Sleep(100 * time.Millisecond)
		clock.Tick()
		processEvents = append(processEvents, Event{process: id, internal_counter: clock.counter, kind: "internal"})
	}

	ts := clock.Send()
	chOut <- Message{sender: id, ts: ts}
	processEvents = append(processEvents, Event{process: id, internal_counter: clock.counter, kind: "send"})

	msg := <-chInt
	clock.Receive(msg.ts)
	processEvents = append(processEvents, Event{process: id, internal_counter: clock.counter, kind: "recv"})
}

func main() {
	var events = make([]Event, 0)
	wg := sync.WaitGroup{}

	ch1 := make(chan Message, 1)
	ch2 := make(chan Message, 1)
	ch3 := make(chan Message, 1)

	wg.Add(3)

	go proces(1, ch1, ch2, events, &wg)
	go proces(2, ch2, ch3, events, &wg)
	go proces(3, ch3, ch1, events, &wg)

	wg.Wait()

	sort.Slice(events, func(i, j int) bool {
		return less(events[i], events[j])
	})
	for i := 0; i < len(events); i++ {
		fmt.Printf("Event details:\n- proces %d\n- internal_counter %d\n- kind: %s", events[i].process, events[i].internal_counter, events[i].kind)
	}
}
