package Adapter

import "fmt"

type Bicicleta struct {
	Marca string
	Color string
}

func (b *Bicicleta) Avanzar(){
	fmt.Println("Avanzando")
}
