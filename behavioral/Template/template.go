package Template

import "fmt"

type WorkerInterface interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	WorkerInterface
}

func NewWorker(w WorkerInterface) *Worker {
	return &Worker{WorkerInterface: w}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct {
}

func (c *Coder) GetUp() {
	fmt.Println("coder GetUp")
}

func (c *Coder) Work() {
	fmt.Println("coder Work")
}

func (c *Coder) Sleep() {
	fmt.Println("coder Sleep")
}
