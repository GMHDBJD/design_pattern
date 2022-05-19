package main

import (
	"errors"
	"fmt"
)

type StateType int

const (
	undefined StateType = iota
	initing
	paused
	running
	stopped
)

type WorkerState interface {
	DoWork()
	ToState(StateType) error
}

type UndefinedWorker struct {
	WorkerState
	stateType StateType
}

func (w *UndefinedWorker) ToState(st StateType) error {
	switch st {
	case initing:
		fmt.Println(w.stateType, "to initing", st)
	default:
		fmt.Println(w.stateType, "undefined to state", st)
		return errors.New("error")
	}
	return nil
}

func (w *UndefinedWorker) DoWork() {
	fmt.Println(w.stateType, "undefined do work")
}

type InitingWorker struct {
	UndefinedWorker
	stateType StateType
}

func (w *InitingWorker) ToState(st StateType) error {
	switch st {
	case initing:
		fmt.Println(w.stateType, "to initing", st)
	case paused:
		fmt.Println(w.stateType, "to paused", st)
	case running:
		fmt.Println(w.stateType, "to running", st)
	case stopped:
		fmt.Println(w.stateType, "to stopped", st)
	default:
		fmt.Println(w.stateType, "undefined to state", st)
		return errors.New("error")
	}
	return nil
}

type PausedWorker struct {
	UndefinedWorker
	stateType StateType
}

func (w *PausedWorker) ToState(st StateType) error {
	switch st {
	case initing:
		fmt.Println(w.stateType, "to initing", st)
	case paused:
		fmt.Println(w.stateType, "to paused", st)
	case running:
		fmt.Println(w.stateType, "to running", st)
	case stopped:
		fmt.Println(w.stateType, "to stopped", st)
	default:
		fmt.Println(w.stateType, "undefined to state", st)
		return errors.New("error")
	}
	return nil
}

type RunningWorker struct {
	UndefinedWorker
	stateType StateType
}

func (w *RunningWorker) ToState(st StateType) error {
	switch st {
	case initing:
		fmt.Println(w.stateType, "to initing", st)
	case paused:
		fmt.Println(w.stateType, "to paused", st)
	case running:
		fmt.Println(w.stateType, "to running", st)
	case stopped:
		fmt.Println(w.stateType, "to stopped", st)
	default:
		fmt.Println(w.stateType, "undefined to state", st)
		return errors.New("error")
	}
	return nil
}

type StoppedWorker struct {
	UndefinedWorker
	stateType StateType
}

func (w *StoppedWorker) ToState(st StateType) error {
	switch st {
	case initing:
		fmt.Println(w.stateType, "to initing", st)
	case paused:
		fmt.Println(w.stateType, "to paused", st)
	case running:
		fmt.Println(w.stateType, "to running", st)
	case stopped:
		fmt.Println(w.stateType, "to stopped", st)
	default:
		fmt.Println(w.stateType, "undefined to state", st)
		return errors.New("error")
	}
	return nil
}

func NewWorkerState(st StateType) WorkerState {
	switch st {
	case initing:
		return &InitingWorker{stateType: initing}
	case paused:
		return &PausedWorker{stateType: paused}
	case running:
		return &RunningWorker{stateType: running}
	case stopped:
		return &StoppedWorker{stateType: stopped}
	default:
		return &UndefinedWorker{stateType: undefined}
	}
}

type Worker struct {
	workerState WorkerState
}

func NewWorker() *Worker {
	return &Worker{workerState: NewWorkerState(undefined)}
}

func (w *Worker) SetState(st StateType) {
	if err := w.workerState.ToState(st); err == nil {
		w.workerState = NewWorkerState(st)
	}
}

func main() {
	w := NewWorker()
	w.SetState(running)
	w.SetState(initing)
	w.SetState(initing)
	w.SetState(running)
	w.SetState(paused)
	w.SetState(stopped)
}
