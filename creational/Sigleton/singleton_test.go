package Sigleton

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	wg:=sync.WaitGroup{}
	wg.Add(200)

	for i:=0;i<100;i++{
		go func() {
			defer wg.Done()
			IncrementAge()
		}()

		go func() {
			defer wg.Done()
			IncrementAge2()
		}()
	}
	wg.Wait()
	p:= GetInstance()
	age:=p.GetAge()
	fmt.Println(age)
}