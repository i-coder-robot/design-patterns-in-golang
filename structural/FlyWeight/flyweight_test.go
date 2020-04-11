package FlyWeight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlyWeightFactory_GetFlyWeight(t *testing.T) {
	factory:=NewFlyWeightFactory()
	hong :=factory.GetFlyWeight("hong beauty")
	xiang :=factory.GetFlyWeight("xiang beauty")

	assert.Len(t,factory.pool,2)
	assert.Equal(t,hong,factory.pool["hong beauty"])
	assert.Equal(t,xiang,factory.pool["xiang beauty"])

}