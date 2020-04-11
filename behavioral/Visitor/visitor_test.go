package Visitor

import "testing"

func TestElement_Accept(t *testing.T) {
	e:=new(Element)
	e.Accept(new(WeiBoVisitor))
	e.Accept(new(IQIYIVisitor))
}
