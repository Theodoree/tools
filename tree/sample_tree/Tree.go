package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Pre         *Node
	Left, Right *Node
	Key         int
	Value       interface{}
}

func NewRoot(key, value int) *Node {
	return &Node{
		Value: value,
		Key:   key,
	}
}

func NewNode(key, Value int) *Node {
	return &Node{
		Value: Value,
		Key:   key,
		Left:  nil,
		Right: nil,
	}
}

func (r *Node) insert(node *Node) {
	current := r
	for {
		if current.Key >= node.Key {
			if current.Left == nil {
				current.Left = node
				node.Pre = current
				break
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = node
				node.Pre = current
				break
			}
			current = current.Right
		}

	}
}
func (r *Node) FindKey(key int) bool {
	current := r
	for {
		if current.Key == key {
			return true
		}
		if current.Key < key && current.Left != nil {
			current = current.Left
		} else if current.Key > key && current.Right != nil {
			current = current.Right
		} else {
			return false
		}
	}
}
func (r *Node) Clear() {
	r.Left = nil
	r.Right = nil
	r.Pre = nil
}

func (r *Node) Delete(key int) bool {
	current := r
	for {
		if key == current.Key {
			Pre := current.Pre
			if Pre.Left == current {
				Pre.Left = current.Right
				current.Clear()
			} else {
				Pre.Right = current.Right
				current.Clear()
			}
			return true
		}
		if current.Key < key && current.Left != nil {
			current = current.Left
		} else if current.Key > key && current.Right != nil {
			current = current.Right
		} else {
			return false
		}
	}

	return false
}

func (r *Node) Min() *Node {
	current := r
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (r *Node) Max() *Node {
	current := r
	for current.Right != nil {
		current = current.Right
	}
	return current
}
func (r *Node) Println() {
	DLR(r, 0)
	LDR(r, 0)
	LRD(r, 0)
}

func DLR(r *Node, high int) {
	space := ""
	for i := 0; i < high; i++ {
		space += "\t"
	}
	fmt.Println(space, "****")
	fmt.Printf("%s ** Key:= %d , Value:=%v  high := %d \n", space, r.Key, r.Value, high)
	fmt.Println(space, "****")
	high ++
	if r.Left != nil {
		DLR(r.Left, high)
	}
	if r.Right != nil {
		DLR(r.Right, high)
	}
}
func LDR(r *Node, high int) {
	high ++
	if r.Left != nil {
		LDR(r.Left, high)
	}
	time.Sleep(time.Second)
	space := ""
	for i := 0; i < high; i++ {
		space += "\t"
	}
	fmt.Println(space, "****")
	fmt.Printf("%s ** Key:= %d , Value:=%v  high := %d \n", space, r.Key, r.Value, high)
	fmt.Println(space, "****")
	if r.Right != nil {
		LDR(r.Right, high)
	}
}

func LRD(r *Node, high int) {
	high ++
	if r.Left != nil {
		LRD(r.Left, high)
	}
	if r.Right != nil {
		LRD(r.Right, high)
	}
	space := ""
	for i := 0; i < high; i++ {
		space += "\t"
	}
	fmt.Println(space, "****")
	fmt.Printf("%s ** Key:= %d , Value:=%v  high := %d \n", space, r.Key, r.Value, high)
	fmt.Println(space, "****")
}

func main() {
	Root := NewRoot(40, 5)
	Root.insert(NewNode(48, 67))
	Root.insert(NewNode(65, 32))
	Root.insert(NewNode(13, 32))
	Root.insert(NewNode(23, 32))
	Root.insert(NewNode(16, 32))
	Root.insert(NewNode(28, 32))
	Root.insert(NewNode(27, 32))
	Root.insert(NewNode(26, 32))
	rand.Seed(time.Now().Unix())
	for i := 100; i < 110; i++ {
		a := rand.Intn(i) + 20
		Root.insert(NewNode(a, i))
	}
	Root.Println()

}
