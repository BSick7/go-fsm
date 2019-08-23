package examples

import (
	"github.com/BSick7/go-fsm"
)

var (
	DoorStateOpen = "open"
	DoorStateClosed = "closed"
	DoorStateOpening = "opening"
	DoorStateClosing = "closing"

	Door = fsm.Machine{
		States: map[string]fsm.State{
			DoorStateClosed: {
				Exits: []string{DoorStateClosing},
				Ignore:  []string{DoorStateOpening},
			},
			DoorStateClosing: {
				Exits: []string{DoorStateClosed},
				Ignore: []string{DoorStateOpening},
			},
			DoorStateOpen: {
				Exits: []string{DoorStateOpening},
				Ignore: []string{DoorStateClosing},
			},
			DoorStateOpening: {
				Exits: []string{DoorStateOpen},
				Ignore: []string{DoorStateClosing},
			},
		},
	}
)
