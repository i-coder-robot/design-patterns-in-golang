package strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategyA:= NewStrategyA()
	c:=NewContext()
	c.SetStrategy(strategyA)
	c.Execute()

	strategyB := NewStrategyB()
	c.SetStrategy(strategyB)
	c.Execute()
}
