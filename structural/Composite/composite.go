package Composite

import "fmt"

type Component interface {
	Traverse()
}

type Leaf struct {
	value int
}

func NewLeaf(value int) *Leaf {
	return &Leaf{value}
}

func (l *Leaf) Traverse() {
	fmt.Println(l.value)
}

type Composite struct {
	children []Component
}

func NewComposite() *Composite {
	return &Composite{children: make([]Component, 0)}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Traverse() {
	for idx, _ := range c.children {
		c.children[idx].Traverse()
	}
}
