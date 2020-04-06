package Responsibility_Chain

import (
	"fmt"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	wang := NewHandler("lao wang", nil, 1)
	zhang := NewHandler("lao zhang", wang, 2)

	r := wang.Handler(1)
	fmt.Println(r)
	r = zhang.Handler(2)
	fmt.Println(r)

}
