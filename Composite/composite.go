package Composite

type Item interface {
	Add(Item)
	String() string
	Price() int
}
