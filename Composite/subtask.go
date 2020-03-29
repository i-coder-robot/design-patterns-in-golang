package Composite

import (
	"fmt"
	"strings"
)

type SubTask struct {
	name string
	price int
}

func (s *SubTask) Add(i Item)  {
	fmt.Println("You don't accept subtrareas")
}

func (s *SubTask) String() string{
	sb := strings.Builder{}
	sb.WriteString("\t\t|-- ")
	sb.WriteString(s.name)
	sb.WriteString("\n")

	return sb.String()
}

func (s *SubTask) Price() int {
	return s.price
}