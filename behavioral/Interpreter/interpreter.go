package Interpreter

import "strings"

type Expression interface {
	Interpret(variables map[string]Expression) int
}

type Integer struct {
	integer int
}

func (n *Integer) Interpret(variables map[string]Expression) int{
	return n.integer
}

type Plus struct {
	leftOperand Expression
	rightOperand Expression
}

func (p *Plus)  Interpret(variables map[string]Expression) int{
	return p.leftOperand.Interpret(variables) + p.rightOperand.Interpret(variables)
}

func (e *Evaluator) Interpret(context map[string]Expression) int  {
	return e.syntaxTree.Interpret(context)
}

type Variable struct {
	name string
}

type Node struct {
	value interface{}
	next *Node
}

type Stack struct {
	top *Node
	size int
}

func (s *Stack) Push(value interface{}) {
	s.top= &Node{
		value: value,
		next:  s.top,
	}
	s.size++
}

func (v *Variable)  Interpret(variables map[string]Expression) int{
	value,found := variables[v.name]
	if !found{
		return 0
	}
	return value.Interpret(variables)
}

func (s *Stack) Pop() interface{}  {
	if s.size==0{
		return nil
	}
	value:= s.top.value
	s.top=s.top.next
	s.size--
	return value
}

type Evaluator struct {
	syntaxTree Expression
}



func NewEvaluator(expression string) *Evaluator{
	expressionStack := new(Stack)
	for _,token := range strings.Split(expression," "){
		switch token {
		case "+":
			right := expressionStack.Pop().(Expression)
			left := expressionStack.Pop().(Expression)
			subExpression := &Plus{left,right}
			expressionStack.Push(subExpression)
		default:
			expressionStack.Push(&Variable{token})
		}
	}
	syntaxTree := expressionStack.Pop().(Expression)
	return &Evaluator{syntaxTree:syntaxTree}
}