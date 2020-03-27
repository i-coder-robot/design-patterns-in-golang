package Observer

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	for x := range Fib(10) {
		fmt.Println(x)
	}

	n := eventSubject{Observers: sync.Map{}}

	obs1 := EventObserver{
		ID:   1,
		Time: time.Now(),
	}

	obs2 := EventObserver{
		ID:   2,
		Time: time.Now(),
	}

	n.AddListener(&obs1)
	n.AddListener(&obs2)

	for x := range Fib(100) {
		n.Notify(Event{Data: x})
	}

}

func test() int{
	return 12
}

