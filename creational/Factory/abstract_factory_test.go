package Factory

import "testing"

func TestNewSimpleShapeFactory(t *testing.T) {
	factory := NewSimpleShapeFactory()
	shape := factory.CreateCurvedShape()
	shape.Draw()

	shape = factory.CreateStraightShape()
	shape.Draw()
}
