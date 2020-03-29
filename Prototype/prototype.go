package Prototype

import "fmt"

type prototype struct {
	name string
	age int
	friends []string
	color *string
	phones map[string]string
}

func (p *prototype) String() string {
	return fmt.Sprintf(
		"Nombre: %s, Edad: %d, Amigos: %v, Color: %s, Teléfonos: %v",
		p.name,
		p.age,
		p.friends,
		*p.color,
		p.phones,
	)
}

func (p *prototype) Clone() prototype{
	friends := make([]string,len(p.friends),len(p.friends))
	copy(friends,p.friends)

	color:= *p.color

	phones:= make(map[string]string)
	for k,v:=range p.phones{
		phones[k]=v
	}

	return prototype{
		name:    p.name,
		age:     p.age,
		friends: friends,
		color:   &color,
		phones:  phones,
	}
}