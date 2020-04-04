package Iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	array := []interface{}{1, 3, 9, 2, 8}
	a := 0
	iterator := ArrayIterator{array: array, index: &a}
	for it := iterator; iterator.HasNext(); iterator.Next() {
		index, value := it.Index(), it.Value().(int)
		if value != array[*index] {
			fmt.Println("error...")
		}
		fmt.Println(*index, value)
	}
}
