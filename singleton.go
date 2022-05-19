package main

import (
	"fmt"
	"sync"
)

type SingleTon struct {
	i int
}

var singleTon *SingleTon
var once sync.Once

func (s *SingleTon) Hello() {
	fmt.Println("hello ", s.i)
	s.i++
}

func GetSingleTon() *SingleTon {
	once.Do(func() {
		singleTon = &SingleTon{i: 0}
	})
	return singleTon
}

func main() {
	s := GetSingleTon()
	s.Hello()
	s.Hello()
	b := GetSingleTon()
	b.Hello()
	s.Hello()
	b.Hello()
}
