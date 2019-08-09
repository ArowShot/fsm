package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arowshot/fsm"
)

func main() {
	basicfsm := fsm.FSM{}

	turtle := turtle{}
	var idleState fsm.State = &IdleState{&turtle}
	var eatingState fsm.State = &EatingState{&turtle}

	basicfsm.AddState("idle", &idleState)
	basicfsm.AddState("eating", &eatingState)

	basicfsm.SetState(&idleState)

	fsmDone := basicfsm.RunMachine()
	defer close(fsmDone)

	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println(err)
			break loop
		}
		switch char {
		case 'F':
			turtle.hunger++
		case 'f':
			turtle.hunger--
		case 'q':
			break loop
		}
	}

	fmt.Println(basicfsm)
}
