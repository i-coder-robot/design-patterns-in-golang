package Mediator

import "fmt"

type Mediator interface {
	Communicate(who string)
}

type WildStallion interface{
	SetMediator(mediator Mediator)
}

type Bill struct {
	mediator Mediator
}

func (b *Bill) SetMediator(mediator Mediator){
	b.mediator=mediator
}

func (b *Bill) Respond(){
	fmt.Println( "Bill: What?")
	b.mediator.Communicate("Bill")
}

type Ted struct {
	mediator Mediator
}

func (t *Ted) Talk(){
	fmt.Println("Ted: Bill?")
	t.mediator.Communicate("Ted")
}

func (t *Ted) SetMediator(mediator Mediator){
	t.mediator=mediator
}

func(t *Ted) Respond(){
	fmt.Println("Ted: Strange things are afoot at the Circle K.")
}

type ConcreteMediator struct {
	Bill
	Ted
}

func NewMediator() *ConcreteMediator{
	mediator:=&ConcreteMediator{}
	mediator.Bill.SetMediator(mediator)
	mediator.Ted.SetMediator(mediator)
	return mediator
}

func (m *ConcreteMediator) Communicate(who string)  {
	if who=="Ted"{
		m.Bill.Respond()
	}else{
		m.Ted.Respond()
	}
}


