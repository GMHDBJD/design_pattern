package main

import "fmt"

type CommandType int

const (
	CommandAType CommandType = iota
	CommandBType
)

type Command interface {
	Execute()
	GetType() CommandType
}

type CommandA struct {
}

func (c *CommandA) Execute() {
	fmt.Println("command a")
}

func (c *CommandA) GetType() CommandType {
	return CommandAType
}

type CommandB struct {
}

func (c *CommandB) GetType() CommandType {
	return CommandBType
}

func (c *CommandB) Execute() {
	fmt.Println("command b")
}

type Dispatcher struct {
	commands map[CommandType]Command
}

func NewDispatcher()*Dispatcher{
	return &Dispatcher{
		commands: map[CommandType]Command{},
	}
}

func (d *Dispatcher) AddCommand(c Command) {
	d.commands[c.GetType()] = c
}

func (d *Dispatcher) DelCommand(c Command) {
	delete(d.commands, c.GetType())
}

func (d *Dispatcher) RunCommand(t CommandType) {
	if c, ok := d.commands[t]; ok {
		c.Execute()
	}
}

func main() {
	ca := &CommandA{}
	cb := &CommandB{}
	d := NewDispatcher()
	d.AddCommand(ca)
	d.AddCommand(cb)
	d.RunCommand(CommandAType)
	d.RunCommand(CommandBType)
	d.DelCommand(ca)
	d.DelCommand(cb)
	d.RunCommand(CommandAType)
	d.RunCommand(CommandBType)
}
