package Composite

import (
	"fmt"
	"testing"
)

func TestComposite_Traverse(t *testing.T) {
	containers := make([]Composite, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			containers[i].Add(NewLeaf(i*3 + j))
		}
	}
	for i := 1; i < 4; i++ {
		containers[0].Add(&(containers[i]))
	}
	for i := 0; i < 4; i++ {
		containers[i].Traverse()
		fmt.Printf("\n")
	}
}
