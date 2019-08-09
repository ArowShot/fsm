package fsm

// FSM is a finite state machine
type FSM struct {
	states        map[string]*State
	currentState  *State
	previousState *State // Keep track of the previous state so we know when to call the OnEnter function
}

// GetStateByName returns a pointer to a state given it's name
//
// Returns nil if state is not found
func (fsm *FSM) GetStateByName(name string) (state *State) {
	return fsm.states[name] // Returns nil if not found in map
}

// SetState will set the fsm's current state to a given state
//
// Returns nil if state is not found
func (fsm *FSM) SetState(state *State) {
	fsm.currentState = state
}

// AddState adds a state to the finite state machine if it does not already exist and they name is not ""
//
// Returns true if added successfully
func (fsm *FSM) AddState(name string, state *State) (okay bool) {
	if fsm.states == nil { // If there is no states map then create one
		fsm.states = make(map[string]*State)
	}

	_, exists := fsm.states[name] // Second return is true if a key exists in a map
	if exists || name == "" {
		return false // Return false if it already exists or name is ""
	}

	fsm.states[name] = state
	return true // Otherwise add the state and return true
}

// RunMachine will start the finite state machine
//
// To stop the machine close the returned channel
//
// This method should only be called once per machine
func (fsm *FSM) RunMachine() (done chan struct{}) {
	if fsm.states == nil { // If there is no states map then create one
		fsm.states = make(map[string]*State)
	}
	done = make(chan struct{}) // Make a channel with an empty struct

	go func() { // Start a new thread to run the machine
		for { // Loop infinitely
			select {
			case <-done: // If the done channel is closed stop the infinite loop
				return
			default: // Otherwise run the state actions
				if fsm.currentState != nil { // Only try to run the state if there is one to begin with
					if fsm.currentState != fsm.previousState { // If the prevouos state was different, then the current state is new and the OnEnter method should be called
						(*fsm.currentState).OnEnter(fsm)
						fsm.previousState = fsm.currentState // Update the previous state
					} else { // Otherwise call the OnStay method
						nextStateName := (*fsm.currentState).OnStay(fsm)
						if nextStateName != "" { // If the nextStateName isn't empty
							nextState := fsm.GetStateByName(nextStateName)
							if nextState != nil { // If the new state exists
								fsm.currentState = nextState // Set the current state as the new state
							}
						}
					}
				}
			}
		}
	}()

	return done // Return the stopping channel
}
