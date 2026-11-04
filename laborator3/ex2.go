// 		P1
// 	 / 		\
// 	|		 P2
// 	 \		/
// 		P3

package main

type LamportClock struct{ counter int }

// incrementeaza contorul
func (l *LamportClock) Tick() {
	l.counter++
}
func (l *LamportClock) Send()          {}
func (l *LamportClock) Receive(ts int) {}

type Message struct{ sender, ts int }

type Event struct {
	process          int
	internal_counter int
	kind             string // " internal ", "send", "recv"
}

// TODO: retuneaza true daca a precede b
func less(a, b Event) bool { return false }
