package main

import "fmt"

type Observer interface {
	ID() string
	Update(interface{})
}

type Subject interface {
	AddObserver(Observer)
	DelObserver(Observer)
	NotifyObserver(interface{})
	hasChanged() bool
	Update()
}

type ObserverA struct {
}

func (o *ObserverA) Update(i interface{}) {
	num, ok := i.(int)
	if ok {
		fmt.Println(num)
	}
}

func (o *ObserverA) ID() string {
	return "A"
}

type ObserverB struct {
}

func (o *ObserverB) Update(i interface{}) {
	num, ok := i.(int)
	if ok {
		fmt.Println(num + 1)
	}
}

func (o *ObserverB) ID() string {
	return "B"
}

type SubjectA struct {
	i         int
	changed   bool
	observers map[string]Observer
}

func NewSubjectA(i int) *SubjectA {
	return &SubjectA{
		i:         i,
		observers: make(map[string]Observer),
	}
}

func (s *SubjectA) AddObserver(o Observer) {
	s.observers[o.ID()] = o
}

func (s *SubjectA) DelObserver(o Observer) {
	delete(s.observers, o.ID())
}

func (s *SubjectA) NotifyObservers() {
	if s.hasChanged() {
		for _, o := range s.observers {
			o.Update(s.i)
		}
		s.setChanged(false)
	}
}

func (s *SubjectA) hasChanged() bool {
	return s.changed
}

func (s *SubjectA) setChanged(b bool) {
	s.changed = b
}

func (s *SubjectA) Update() {
	s.i++
	s.setChanged(true)
}

func main() {
	subject := NewSubjectA(0)
	oA := &ObserverA{}
	oB := &ObserverB{}
	subject.AddObserver(oA)
	subject.AddObserver(oB)

	subject.NotifyObservers()

	subject.Update()
	subject.NotifyObservers()

	subject.NotifyObservers()

	subject.Update()
	subject.Update()
	subject.NotifyObservers()

	subject.DelObserver(oB)
	subject.Update()
	subject.NotifyObservers()
}
