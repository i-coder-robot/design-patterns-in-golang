package Adapter

func Factory(s string) Adapter{
	switch s {
	case "automovil":
		d:=Automovil{}
		return &AutomovilAdapter{&d}
	case "bicicleta":
		d:= Bicicleta{}
		return &BicicletaAdapter{&d}
	}
	return nil
}