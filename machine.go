package fsm

type Machine struct {
	States map[string]State
}

type State struct {
	Exits  []string
	Ignore []string
}
