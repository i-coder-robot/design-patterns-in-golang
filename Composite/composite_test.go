package Composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	p := Project{name: "Upload images"}
	t1 := Task{name: "Mockup", responsable: "UI Designer", price: 1000}

	t2 := Task{name: "Markup", responsable: "Web Designer"}
	st21 := SubTask{name: "HTML", price: 300}
	st22 := SubTask{name: "CSS", price: 700}
	t2.Add(&st21)
	t2.Add(&st22)

	t3 := Task{name: "JS", responsable: "Frontend Developer", price: 1000}

	t4 := Task{name: "API Backend", responsable: "Backend Developer"}
	st41 := SubTask{name: "Authentication", price: 100}
	st42 := SubTask{name: "DB connection", price: 900}
	t4.Add(&st41)
	t4.Add(&st42)

	t5 := Task{name: "Database", responsable: "DBA", price: 1000}

	p.Add(&t1)
	p.Add(&t2)
	p.Add(&t3)
	p.Add(&t4)
	p.Add(&t5)

	fmt.Println(&p)
}
