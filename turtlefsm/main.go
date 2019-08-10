package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arowshot/fsm"
)

func main() {
	// Create the FSM
	basicfsm := fsm.FSM{}

	// Create the turtle
	turtle := turtle{}
	var idleState fsm.State = &IdleState{&turtle}     // Create idle state with a refernece to the turtle
	var eatingState fsm.State = &EatingState{&turtle} // Create eating state with a refernece to the turtle
	var warmingState fsm.State = &IdleState{&turtle}  // Create warming state with a refernece to the turtle

	// Add states to fsm
	basicfsm.AddState("idle", &idleState)
	basicfsm.AddState("eating", &eatingState)
	basicfsm.AddState("warming", &warmingState)

	// Set the starting fsm state
	basicfsm.SetState(&idleState)

	// Run the fsm
	fsmDone := basicfsm.RunMachine()
	defer close(fsmDone) // Stop the fsm when the program ends

	reader := bufio.NewReader(os.Stdin) // Create a reader for user input

loop:
	for {
		char, _, err := reader.ReadRune() // Read user input
		if err != nil {
			fmt.Println(err)
			break loop // Quit if there's an error reading
		}
		switch char { // Perform an action on the turtle depending on the user input
		case 'H':
			turtle.hunger++
		case 'h':
			turtle.hunger--
		case 'S':
			turtle.sunlight++
		case 's':
			turtle.sunlight--
		case 'q':
			break loop
		}
	}

	fmt.Println("Turtle FSM done -- have a nice day")
}
