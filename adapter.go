package main

import "fmt"

type Duck interface {
	Guagua()
}

type Bird struct {
}

func (b *Bird) Jiji() {
	fmt.Println("jiji")
}

type DuckBird struct {
	Duck
	b *Bird
}

func (d *DuckBird) Guagua() {
	d.b.Jiji()
}

func callDuck(d Duck) {
	d.Guagua()
}

func main() {
	db := &DuckBird{b: &Bird{}}
	db.Guagua()
	callDuck(db)
}
