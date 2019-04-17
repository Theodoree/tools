package datatype

import (
	"fmt"
	"qiniupkg.com/x/errors.v7"
)

type Linked4 struct {
	First *Node4
	Last  *Node4
}

type Node4 struct {
	Parent *Node4
	Next   *Node4
	Key    int
	Value  int
}

func (l Linked4) NewLinked() *Linked4 {
	return &Linked4{}
}
func (l Linked4) NewNode(key, value int) *Node4 {
	return &Node4{Key: key, Value: value}
}

//从尾部插入O(1)
func (l *Linked4) Insert(key, value int) {
	node := Linked4{}.NewNode(key, value)
	node.Next = l.First
	if l.First == nil {
		l.Last = node
		l.First = node
		node.Parent = l.First
		return
	}
	current := l.Last
	current.Next = node
	node.Parent = current
	l.Last = node
	return
}


//O(1)
func (l *Linked4) DeleteFirst() {
	if l.First == nil {
		return
	}

	l.First = l.First.Next
	l.Last.Next = l.First
	return
}

//O(n)
func (l *Linked4) Delete(key int) bool {
	if l.First == nil {
		return false
	}

	current := l.First
	for current.Key != key {
		if current == nil || current.Next == l.First {
			return false
		}
		current = current.Next
	}
	parent := current.Parent
	next := parent.Next.Next
	parent.Next = next
	return true
}

//O(n)
func (l *Linked4) Find(key int) (int, error) {

	current := l.First
	for current.Next != nil && current.Next != l.First{
		if current.Key != key {
			current = current.Next
		} else {
			return current.Value, nil
		}
	}
	return 0, errors.New("Not Find")

}

func (l *Linked4) Print() {
	if l.First == nil{
		return
	}

	current := l.First

	for current.Next != nil && current.Next != l.First {
		fmt.Println(current.Key, "--", current.Value)
		current = current.Next
	}
	fmt.Println(current.Key, "--", current.Value)

}
