package fsm

// State is an individual state in a finite state machine
//
// OnEnter is called when the state begins
//
// OnStay is called in a loop until the state changes again and returns the next state name
type State interface {
	OnEnter(*FSM)
	OnStay(*FSM) string
}
