package Adapter

import "fmt"

type Target interface {
	Execute()
}

type Adaptee struct {

}

func (a *Adaptee) SpecificExecute() {
	fmt.Println("充电...")
}

type Adapter struct{
	*Adaptee
}

func (a *Adapter) Execute(){
	a.SpecificExecute()
}