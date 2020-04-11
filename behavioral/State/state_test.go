package State

import "testing"

func TestState(t *testing.T) {
	machine := NewMachine()
	machine.Off()
	machine.On()
	machine.On()
	machine.Off()
}