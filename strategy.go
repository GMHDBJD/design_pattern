package main

import "fmt"

type strategyType int

const (
	strategyA strategyType = iota
	strategyB
)

type Strategy interface {
	strategy()
}

type StrategyA struct {
}

func (s *StrategyA) strategy() {
	fmt.Println("a")
}

type StrategyB struct {
}

func (s *StrategyB) strategy() {
	fmt.Println("b")
}

type Worker struct {
	workerStrategy Strategy
}

func NewWorker() *Worker {
	return &Worker{
		workerStrategy: &StrategyA{},
	}
}

func (w *Worker) setStrategy(st strategyType) {
	switch st {
	case strategyA:
		w.workerStrategy = &StrategyA{}
	case strategyB:
		w.workerStrategy = &StrategyB{}
	}
}

func (w *Worker) doWork() {
	w.workerStrategy.strategy()
}

func main() {
	worker := NewWorker()
	worker.doWork()
	worker.setStrategy(strategyB)
	worker.doWork()
}
