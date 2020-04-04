package Strategy

import "fmt"

type Strategy interface {
	Execute()
}

type strategyA struct {
}

func (s *strategyA) Execute() {
	fmt.Println("A plan executed.")
}

func NewStrategyA() Strategy {
	return &strategyA{}
}

type strategyB struct {
}

func (s *strategyB) Execute() {
	fmt.Println("B plan executed.")
}

func NewStrategyB() Strategy {
	return &strategyB{}
}

type Context struct {
	strategy Strategy
}

func NewContext() *Context {
	return &Context{}
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}
func (c *Context) Execute() {
	c.strategy.Execute()
}
