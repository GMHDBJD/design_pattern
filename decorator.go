package main

import "fmt"

type Work interface {
	doFunc()
}

type Worker struct{}

type Decorator struct {
	Work
}

func NewWorker() *Worker {
	return &Worker{}
}

func NewDecorator(w Work) *Decorator {
	return &Decorator{Work: w}
}

func (*Worker) doFunc() {
	fmt.Println("doFunc")
}

func (d *Decorator) doFunc() {
	fmt.Println("enter")
	d.Work.doFunc()
	fmt.Println("leave")
}

func main() {
	d := NewDecorator(NewWorker())
	d.doFunc()
}
