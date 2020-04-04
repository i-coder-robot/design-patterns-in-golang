package Abstract_Factory

import "testing"

func TestNewSimpleShapeFactory(t *testing.T) {
	factory := NewSimpleShapeFactory()
	food := factory.CreateFood()
	food.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
