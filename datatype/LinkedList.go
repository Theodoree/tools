package datatype

import (
	"errors"
	"fmt"
)

type Linked struct {
	Node *Node
}

type Node struct {
	Next  *Node
	Key   int
	Value int
}

func NewLinked() *Linked {
	return &Linked{}
}
func NewNode(key, value int) *Node {
	return &Node{Key: key, Value: value}
}
func (l *Linked) InsertFirst(key, value int) {
	current := l.Node
	node := NewNode(key, value)

	if current == nil {
		l.Node = NewNode(key, value)
		return
	}
	current = l.Node
	node.Next = current
	l.Node = node
	return

}

func (l *Linked) Insert(key, value int) {
	node := NewNode(key, value)
	if l.Node == nil {
		l.Node = node
		return
	}

	current := l.Node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node

	return
}

func (l *Linked) DeleteFirst() {
	if l.Node == nil {
		return
	}

	node := l.Node.Next
	l.Node = node
	return
}

func (l *Linked) Delete(key int) bool {
	if l.Node == nil {
		return false
	}

	current := l.Node
	parent := current
	for current.Key != key {
		if current == nil {
			return false
		}
		parent = current
		current = current.Next
	}
	next:= parent.Next.Next
	parent.Next = next
	return  true
}

func (l *Linked) Find(key int) (int, error) {

	current := l.Node
	for current.Next != nil {
		if current.Key != key {
			current = current.Next
		} else {
			return current.Value, nil
		}
	}

	return 0, errors.New("Not Find")

}

func (l *Linked) Print() {
	current := l.Node

	for current.Next != nil {
		fmt.Println(current.Key, "--", current.Value)
		current = current.Next
	}
	fmt.Println(current.Key, "--", current.Value)

}
