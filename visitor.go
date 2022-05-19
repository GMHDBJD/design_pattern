package main

import "fmt"

type Visitor interface {
	VisitA(Object)
	VisitB(Object)
	VisitC(Object)
}

type VisitorFirst struct {
	Visitor
}

type VisitorSecond struct {
	Visitor
}

type Object interface {
	Accept(Visitor)
	GetName() string
}

type A struct {
	Object
}

type B struct {
	Object
}

type C struct {
	Object
}

func (o *A) Accept(v Visitor) {
	v.VisitA(o)
}

func (o *A) GetName() string {
	return "a"
}

func (o *B) Accept(v Visitor) {
	v.VisitB(o)
}

func (o *B) GetName() string {
	return "b"
}

func (o *C) Accept(v Visitor) {
	v.VisitC(o)
}

func (o *C) GetName() string {
	return "c"
}

func (v *VisitorFirst) VisitA(o Object) {
	fmt.Println("First", o.GetName())
}

func (v *VisitorFirst) VisitB(o Object) {
	fmt.Println("First", o.GetName())
}

func (v *VisitorFirst) VisitC(o Object) {
	fmt.Println("First", o.GetName())
}

func (v *VisitorSecond) VisitA(o Object) {
	fmt.Println("Second", o.GetName())
}

func (v *VisitorSecond) VisitB(o Object) {
	fmt.Println("Second", o.GetName())
}

func (v *VisitorSecond) VisitC(o Object) {
	fmt.Println("Second", o.GetName())
}

func main() {
	vf := &VisitorFirst{}
	vs := &VisitorSecond{}
	obs := []Object{&A{}, &B{}, &C{}}
	for _, ob := range obs {
		ob.Accept(vf)
		ob.Accept(vs)
	}
}
