package Composite

import (
	"strconv"
	"strings"
)

type Project struct {
	name string
	tasks []Item
}

func (p *Project) Add(i Item){
	p.tasks = append(p.tasks, i)
}

func (p *Project) String() string {
	sb := strings.Builder{}
	sb.WriteString(p.name)
	sb.WriteString(" $")
	sb.WriteString(strconv.Itoa(p.Price()))
	sb.WriteString("\n")
	for _, v := range p.tasks {
		sb.WriteString(v.String())
	}

	return sb.String()
}

func (p *Project) Price() int {
	price :=0
	for _,v:=range p.tasks{
		price+=v.Price()
	}
	return price
}
