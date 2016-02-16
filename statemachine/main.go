package main

import "fmt"

type stateMachine struct {
	tr chan int
	st chan string
}

func newStateMachine() *stateMachine {
	sm := &stateMachine{
		tr: make(chan int),
		st: make(chan string),
	}
	go sm.loop()
	return sm
}

func (sm *stateMachine) loop() {
	state := "A"
	for {
		select {
		case i := <-sm.tr:
			fmt.Printf("state %s + %d => ", state, i)
			switch {
			case state == "A" && i == 0:
				state = "A"
			case state == "A" && i == 1:
				state = "B"
			case state == "B" && i == 0:
				state = "C"
			case state == "B" && i == 1:
				state = "A"
			case state == "C" && i == 0:
				state = "B"
			case state == "C" && i == 1:
				state = "A"
			}
			fmt.Printf("state %s\n", state)
		case sm.st <- "state " + state:
		}
	}
}

func (sm *stateMachine) send(i int)    { sm.tr <- i }
func (sm *stateMachine) state() string { return <-sm.st }

func main() {
	sm := newStateMachine()
	sm.send(1)          // "state A + 1 => state B"
	sm.send(0)          // "state B + 0 => state C"
	println(sm.state()) // "state C"
}
