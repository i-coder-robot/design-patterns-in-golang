package Builder

type Builder interface {
	Build()
}

type Director struct {
	builder Builder
}

func NewDirector(b Builder) Director {
	return Director{builder: b}
}

func (d *Director) Construct() {
	d.builder.Build()
}

type ConcreteBuilder struct {
	built bool
}

func NewConcreteBuilder() ConcreteBuilder {
	return ConcreteBuilder{false}
}

func (b *ConcreteBuilder) Build() {
	b.built = true
}

type Product struct {
	Built bool
}

func (b *ConcreteBuilder) GetResult() Product {
	return Product{b.built}
}
