package Visitor

import (
	"fmt"
)

type IVisitor interface {
	Visit()
}

type WeiboVisitor struct {
}

func (w WeiboVisitor) Visit() {
	fmt.Println("这是微博")
}

type IQiYiVisitor struct {
}

func (i IQiYiVisitor) Visit() {
	fmt.Println("这里是爱奇艺")
}

type IElement interface {
	Accept(visitor IVisitor)
}

type Element struct {
}

func (el Element) Accept(v IVisitor) {
	v.Visit()
}
