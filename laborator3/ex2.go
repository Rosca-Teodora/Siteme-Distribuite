package main

type LamportClock struct{ c int }

func (l *LamportClock) Tick() int { /* TODO */ return 0 }

func (l *LamportClock) Send() int { /* TODO */ return 0 }

func (l *LamportClock) Receive(ts int) int { /* TODO */ return 0 }

type Event struct{ Proc, C int }

// TODO: true daca a precede b in ordinea totala
func less(a, b Event) bool { return false }
