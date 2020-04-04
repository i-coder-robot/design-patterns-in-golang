package Factory

import "fmt"

type Shape interface {
	Draw()
}

type circle struct {
}

func (c *circle) Draw() {
	fmt.Println("it is a circle.")
}

type square struct {
}

func (c *square) Draw() {
	fmt.Println("it is square.")
}

type ellipse struct {
}

func (e *ellipse) Draw() {
	fmt.Println("it is ellipse")
}

type ShapeFactory interface {
	CreateCurvedShape() Shape
	CreateStraightShape() Shape
}

type simpleShapeFactory struct {
}

func NewSimpleShapeFactory() ShapeFactory {
	return &simpleShapeFactory{}
}

func (s *simpleShapeFactory) CreateCurvedShape() Shape {
	return &circle{}
}

func (s *simpleShapeFactory) CreateStraightShape() Shape {
	return &square{}
}
