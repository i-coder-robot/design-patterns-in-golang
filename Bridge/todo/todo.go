package todo

import "github.com/i-coder-robot/design-patterns-in-golang/Bridge/list"

type ToDo interface {
	SetList(l list.List)
	Add(name string)
	Print()
}
