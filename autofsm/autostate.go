package autofsm

import "github.com/arowshot/fsm"

// transition has a string for the next state and a condition to check before
// changing to that state
type transition struct {
	NextState string
	Condition func() bool
}

// AutoState is a type of State which automatically transitions to another state
// when certain conditions are met
type AutoState struct {
	Action      func(*fsm.FSM)
	transitions []transition
}

// AddTransition adds a transition to another state based on a condition
func (as *AutoState) AddTransition(nextState string, condition func() bool) {
	as.transitions = append(as.transitions, transition{nextState, condition})
}

// OnEnter is called when the state begins
func (as *AutoState) OnEnter(fsm *fsm.FSM) {
	// Don't do anything special, just get to the OnStay
}

// OnStay is called on a loop until a trnasition condition os fulfilled
func (as *AutoState) OnStay(fsm *fsm.FSM) (nextState string) {
	as.Action(fsm) // Run the action code

	for _, trans := range as.transitions { // Loop through all the transitions
		if trans.Condition == nil || trans.Condition() { // Check the condition on each transition if there is one
			return trans.NextState // Go to the next state if the condition is true or there is no condition
		}
	}

	return "" // Return "" which will not change the state
}

// ToState returns the AutoState as a State type
func (as *AutoState) ToState() *fsm.State {
	var state fsm.State = as
	return &state
}
