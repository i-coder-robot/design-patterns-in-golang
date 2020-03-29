package cmd

import (
	"github.com/i-coder-robot/design-patterns-in-golang/Bridge/list"
	"github.com/i-coder-robot/design-patterns-in-golang/Bridge/todo"
	"testing"
)

func Test(t *testing.T)  {
	myToDos := factoryToDo("any")
	rendering :=factoryList("bullet")
	myToDos.SetList(rendering)

	myToDos.Add("study")
	myToDos.Add("sports")
	myToDos.Add("work")
	myToDos.Add("happy")

	myToDos.Print()
}


func factoryToDo(s string) todo.ToDo{
	if s=="unique"{
		return todo.NewUnique()
	}
	return todo.NewAny()
}

func factoryList(s string) list.List{
	if s == "plain"{
		return list.NewPlain()
	}
	return list.NewBullet('*')
}