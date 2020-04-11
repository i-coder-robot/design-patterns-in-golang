package Memento

import "testing"

func TestMemento(t *testing.T) {
	n:= NewNUmber(10)
	n.Double()
	memento:=n.CreateMemento()
	n.Half()
	n.ReinstateMemento(memento)
}