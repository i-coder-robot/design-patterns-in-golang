package Composite

import (
	"strconv"
	"strings"
)

type Task struct {
	name string
	responsable string
	price int
	subTasks []Item
}

func (t *Task) Add(i Item){
	t.subTasks = append(t.subTasks,i)
}

func (t *Task) String() string  {
	sb:=strings.Builder{}
	sb.WriteString("\t|--")
	sb.WriteString(t.name)
	sb.WriteString(" - ")
	sb.WriteString(t.responsable)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(t.Price()))
	sb.WriteString("\n")
	for _,v := range t.subTasks{
		sb.WriteString(v.String())
	}
	return sb.String()
}

func (t *Task) Price() int {
	price:= t.price
	for _,v:=range t.subTasks{
		price +=v.Price()
	}
	return price
}