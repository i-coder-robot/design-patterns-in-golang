package Sigleton

func IncrementAge() {
	p:= GetInstance()
	p.IncrementAge()
}
