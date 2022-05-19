package main

import "fmt"

type Handler interface {
	Handle(int)
}

type DefaultHanlder struct {
	Handler
	next Handler
}

func (h *DefaultHanlder) SetNext(hd Handler) {
	h.next = hd
}

type HandlerA struct {
	DefaultHanlder
}

type HandlerB struct {
	DefaultHanlder
}

func (h *HandlerA) Handle(i int) {
	if i < 10 {
		fmt.Println("handleA", i)
	} else if h.next != nil {
		h.next.Handle(i)
	}
}

func (h *HandlerB) Handle(i int) {
	if i < 20 {
		fmt.Println("handleB", i)
	} else if h.next != nil {
		h.next.Handle(i)
	}
}

func main() {
	ha := &HandlerA{}
	hb := &HandlerB{}
	ha.SetNext(hb)
	ha.Handle(1)
	ha.Handle(11)
	ha.Handle(21)
}
