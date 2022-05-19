package main

import "fmt"

type Duck struct{
}

func (d*Duck)Guagua(){
	fmt.Println("guagua")
}

func (d*Duck)Walk(){
	fmt.Println("walk")
}

type MyDuck struct{
	d *Duck
}

func (m*MyDuck)WalkAndGua(){
	m.d.Walk()
	m.d.Guagua()
}

func main(){
	myDuck:=&MyDuck{d: &Duck{}}
	myDuck.WalkAndGua()
}