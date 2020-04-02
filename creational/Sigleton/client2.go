package Sigleton

func IncrementAge2() {
	p:= GetInstance()
	p.IncrementAge()
}