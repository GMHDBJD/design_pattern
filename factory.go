package main

import "fmt"

type Work interface {
	doFunc()
}

type Worker struct {
}

func (w *Worker) doFunc() {
	fmt.Println("doFunc")
}

type WorkerFactory struct {
}

func (*WorkerFactory) CreateWorker() Work {
	return &Worker{}
}

func main() {
	workerFactory := &WorkerFactory{}
	work := workerFactory.CreateWorker()
	work.doFunc()
}
