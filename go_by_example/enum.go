package main

import "fmt"

// Type alias
type ServerState int

// Enumerated types (enums) are idiom,
// defined as constants
const (
	// `iota`: Successive untyped integer constant (0, 1, 2, ...)
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// String expressions
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// Implement fmt.Stringer inteface for print
func (ss ServerState) String() string {
	return stateName[ss]
}

// Emulates state transition for a server
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	ns3 := transition(ns2)
	fmt.Println(ns3)
}
