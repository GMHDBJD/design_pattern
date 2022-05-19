package main

import "fmt"

type Computer interface {
	Print()
}

type Printer interface {
	Print()
}

type HPPrinter struct {
	Printer
}

type SonyPrinter struct {
	Printer
}

func (p *HPPrinter) Print() {
	fmt.Println("HP")
}

func (p *SonyPrinter) Print() {
	fmt.Println("sony")
}

type Windows struct {
	Computer
	P Printer
}

type Mac struct {
	Computer
	P Printer
}

func (c *Windows) Print() {
	fmt.Println("windows")
	c.P.Print()
}

func (c *Mac) Print() {
	fmt.Println("mac")
	c.P.Print()
}

func main() {
	mac := &Mac{P:&HPPrinter{}}
	mac.Print()
	mac.P=&SonyPrinter{}
	mac.Print()
	win := &Windows{P:&HPPrinter{}}
	win.Print()
	win.P=&SonyPrinter{}
	win.Print()
}
