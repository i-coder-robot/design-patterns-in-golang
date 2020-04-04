package Strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategy := NewStrategyB()
	context := NewContext()
	context.SetStrategy(strategy)
	context.Execute()
}
