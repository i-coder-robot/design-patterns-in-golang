package Adapter

type Adapter interface {
	Mover()
}

type AutomovilAdapter struct {
	Auto *Automovil
}

func (a *AutomovilAdapter) Mover() {
	if !a.Auto.Encendido{
		a.Auto.Encender()
	}
	a.Auto.Acelerar()
}

type BicicletaAdapter struct {
	Bici *Bicicleta
}

func (b *BicicletaAdapter) Mover(){
	b.Bici.Avanzar()
}