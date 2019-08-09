package main

import (
	"fmt"
	"time"

	"github.com/arowshot/fsm"
)

type turtle struct {
	hunger   int
	sunlight int
}

// IdleState is the turtle's idle state
type IdleState struct {
	turtle *turtle // A reference to the turtle
}

// OnEnter is called when the turtle enters the idle state
func (is *IdleState) OnEnter(fsm *fsm.FSM) {
	fmt.Println("The turtle is entering the idle state")
}

// OnStay is called constantly while the turtle is in the idle state
func (is *IdleState) OnStay(fsm *fsm.FSM) (nextState string) {
	// Output some info
	fmt.Println("The turtle is waiting for something interesting to happen.")
	fmt.Printf("Current hunger: %v, Current sunlight: %v\n", is.turtle.hunger, is.turtle.sunlight)

	if is.turtle.hunger > 5 { // Check turtle's hunger
		return "eating" // Change state to "eating" if the turtle has more than 5 hunger
	}

	if is.turtle.sunlight > 10 { // Check turtle's sunlight
		return "warming" // Change state to "warming" if the turtle has more than 10 sunlight
	}

	fmt.Println("Nothing interesting happened, I'll wait 5 seconds")

	time.Sleep(5 * time.Second) // Wait for 5 seconds

	return "" // Return "" which will not change the state
}

// EatingState is the turtle's eating state
type EatingState struct {
	turtle *turtle // A reference to the turtle
}

// OnEnter is called when the turtle enters the eating state
func (is *EatingState) OnEnter(fsm *fsm.FSM) {
	fmt.Println("The turtle is entering the eating state")
}

// OnStay is called constantly while the turtle is in the eating state
func (is *EatingState) OnStay(fsm *fsm.FSM) (nextState string) {
	fmt.Println("Turtle is eating a food...")

	time.Sleep(2 * time.Second) // Wait 2 seconds
	is.turtle.hunger--          // Then subtract one from hunger to eat a food

	fmt.Printf("Turtle ate a food, now has %v hunger\n", is.turtle.hunger)

	if is.turtle.hunger < 2 { // Check turtle's hunger
		return "idle" // Return to idle if the turtle has eaten enough
	}

	fmt.Println("Turtle is still hungry, eating some more...")

	return "" // Return "" which will not change the state
}
