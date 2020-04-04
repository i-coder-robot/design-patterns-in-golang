package Responsibility_Chain

import (
	"fmt"
	"testing"
)

func TestHandler_Handler(t *testing.T) {
	micheal := NewHandler("Micheal", nil, 1)
	mike := NewHandler("Mike", micheal, 2)
	r := mike.Handler(1)
	fmt.Println(r)
	r = mike.Handler(2)
	fmt.Println(r)
}
