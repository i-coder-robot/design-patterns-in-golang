package Observer

import (
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	//for x:=range Fib(20){
	//	fmt.Println(x)
	//}

	n := eventSubject{Observers: sync.Map{}}

	obs1:= eventObserver{ID: 1,Time:time.Now()}
	obs2:= eventObserver{ID: 2,Time:time.Now()}

	n.AddListener(obs1)
	n.AddListener(obs2)

	for x:=range Fib(10){
		n.Notify(Event{Data: x})
	}
}

