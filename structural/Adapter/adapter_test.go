package Adapter

import "testing"

func TestAdaptee_SpecificExecute(t *testing.T) {
	adapter := Adapter{}
	adapter.Execute()
}