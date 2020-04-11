package Prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) Name() string{
	return p.name
}

func (p *ConcretePrototype) Clone() Prototype{
	return &ConcretePrototype{name:p.name}
}

