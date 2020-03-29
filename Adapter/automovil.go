package Adapter

import "fmt"

type Automovil struct {
	Marca string
	Model uint8
	Encendido bool
}

func (a *Automovil) Encender(){
	if a.Encendido{
		fmt.Println("Ya está encendido")
		return
	}
	fmt.Println("Encendido!")
}

func (a *Automovil) Acelerar(){
	fmt.Println("Acelerando!")
}