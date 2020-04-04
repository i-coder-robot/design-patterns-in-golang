package Iterator

type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
}

type ArrayIterator struct {
	array []interface{}
	index *int
}

func (a *ArrayIterator) Index() *int {

	return a.index
}

func (a *ArrayIterator) Value() interface{} {
	return a.array[*a.index]
}

func (a *ArrayIterator) HasNext() bool {
	return *a.index+1 <= len(a.array)
}

func (a *ArrayIterator) Next() {
	if a.HasNext() {

		*a.index++
	}
}
