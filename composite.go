package main

import "fmt"

type Node interface {
	Travel()
}

type NodeItem struct {
	i int
}

func (n *NodeItem) Travel() {
	fmt.Println(n.i)
}

type Folder struct {
	NodeItem
	i     int
	nodes []Folder
}

func (f *Folder) Travel() {
	fmt.Println(f.i)
	for _, n := range f.nodes {
		n.Travel()
	}
}

func main() {
	f := &Folder{nodes: []Folder{{i: 1}, {i: 2}, {i: 3, nodes: []Folder{{i: 4}, {i: 5}}}}}
	f.Travel()
}
