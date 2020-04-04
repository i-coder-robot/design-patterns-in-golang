package Facade

import "testing"

func TestCarFacade_CreateCompleteCar(t *testing.T) {
	facade := NewCarFacade()
	facade.CreateCompleteCar()
}
