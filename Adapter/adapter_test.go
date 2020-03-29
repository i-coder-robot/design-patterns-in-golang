package Adapter

import (
	"testing"
)

func TestAutomovilAdapter_Mover(t *testing.T) {
	var s string
	//s="automovil"
	s="bicicleta"
	transport :=Factory(s)
	transport.Mover()
}
