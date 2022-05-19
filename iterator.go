package main

import "fmt"

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type DB interface {
	createIterator() Iterator
}

type Item struct {
	i int
}

type ItemIterator struct {
	Iterator
	items []Item
	pos   int
}

func (it *ItemIterator) HasNext() bool {
	return it.pos < len(it.items)
}

func (it *ItemIterator) Next() interface{} {
	item := it.items[it.pos]
	it.pos++
	return item
}

type MyDB struct {
	DB
	items []Item
}

func (db *MyDB) createIterator() ItemIterator {
	return ItemIterator{items: db.items}
}

func main() {
	myDB := &MyDB{items: []Item{{1}, {2}, {3}}}
	iter := myDB.createIterator()
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
