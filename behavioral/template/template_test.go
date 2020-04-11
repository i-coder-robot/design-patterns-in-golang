package template

import "testing"

func TestWorker_Daily(t *testing.T) {
	worker :=NewWorker(&Coder{})

	worker.Daily()
}