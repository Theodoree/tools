package datatype

import (
	"errors"
	"fmt"
)

type Linked2 struct {
	First *Node2
	Last  *Node2
}

type Node2 struct {
	Next  *Node2
	Key   int
	Value int
}

func (l Linked2) NewLinked() *Linked2 {
	return &Linked2{}
}
func (l Linked2) NewNode(key, value int) *Node2 {
	return &Node2{Key: key, Value: value}
}

func (l *Linked2) InsertFirst(key, value int) {
	current := l.First
	node := Linked2{}.NewNode(key, value)

	if current == nil {
		l.First = Linked2{}.NewNode(key, value)
		return
	}
	current = l.First
	node.Next = current
	l.First = node
	return

}

func (l *Linked2) Insert(key, value int) {
	node := Linked2{}.NewNode(key, value)
	node.Next = l.First
	l.Last = node
	if l.First == nil {
		l.First = node
		return
	}

	current := l.First
	for current.Next != nil && current.Next != l.First {
		current = current.Next
	}
	current.Next = node

	return
}

func (l *Linked2) DeleteFirst() {
	if l.First == nil {
		return
	}

	l.First = l.First.Next
	l.Last.Next = l.First
	return
}

func (l *Linked2) Delete(key int) bool {
	if l.First == nil {
		return false
	}

	current := l.First
	parent := current
	for current.Key != key {
		if current == nil || current.Next == l.First {
			return false
		}
		parent = current
		current = current.Next
	}
	next := parent.Next.Next
	parent.Next = next
	return true
}

func (l *Linked2) Find(key int) (int, error) {

	current := l.First
	for current.Next != nil && current.Next != l.First {
		if current.Key != key {
			current = current.Next
		} else {
			return current.Value, nil
		}
	}

	return 0, errors.New("Not Find")

}

func (l *Linked2) Print() {
	current := l.First

	for current.Next != l.First {
		fmt.Println(current.Key, "--", current.Value)
		current = current.Next
	}
	fmt.Println(current.Key, "--", current.Value)

}
