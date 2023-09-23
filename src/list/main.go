package main

import (
	list "list/storage"
)

func main() {
	l := list.NewList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(4)
	l.Add(4)
	l.Add(4)
	l.RemoveByValue(1)
	l.RemoveByIndex(3)
	l.Print()
}
