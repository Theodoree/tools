package binary_serach_tree

import (
	"fmt"
	"time"
)

type Root struct {
	Root  *Node
	Count int
}

type Node struct {
	Right, Left *Node
	Key         int
	Value       int
}

func NewRoot() *Root {
	return &Root{
		&Node{Left: nil, Right: nil}, 0,
	}

}

func NewNode(key int, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

/*
func (r *Root) Insert(root *Node, key int, value string) *Node {

	current := root
	for {
		if current.Key == key {
			current.Value = value
			break;
		} else if current.Key > key {
			if current.Left == nil {
				r.Count++
				current.Left = NewNode(key, value)
				break;
			} else {
				current = current.Left
			}
		} else if current.Key < key{
			if current.Right == nil {
				r.Count++
				current.Right = NewNode(key, value)
				break;
			} else {
				current = current.Right
			}
		}
	}

	return current
}
*/

func (r *Root) Insert(root *Node, key int, value int) *Node {
	if root == nil {
		r.Count++
		return NewNode(key, value)
	}
	if root.Key == key {
		root.Value = value
	} else if root.Key < key {
		root.Right = r.Insert(root.Right, key, value)
	} else if root.Key > key {
		root.Left = r.Insert(root.Left, key, value)
	}
	return root
}

func (r *Root) Contain(key int) bool {
	return contain(r.Root, key)
}
func contain(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key == node.Key {
		return true
	} else if node.Key > key {
		return contain(node.Left, key)
	} else {
		return contain(node.Right, key)
	}
}

func (r *Root) Search(root *Node, key int) *Node {
	return search(root, key)
}
func search(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if root.Key == key {
		return root
	} else if root.Key > key {
		return search(root.Left, key)
	} else {
		return search(root.Right, key)
	}
}

func (r *Root) Delete(target int) bool {
	return delete(r.Root, target)
}
func delete(node *Node, target int) bool {
	if node.Left.Key == target {
		if node.Left.Left != nil {
			node.Left = node.Left.Left
		} else if node.Left.Right != nil {
			node.Left = node.Left.Right
		} else {
			node.Left = nil
		}
		return true
	} else if node.Right.Key == target {
		if node.Right.Left != nil {
			node.Right = node.Right.Left
		} else if node.Right.Right != nil {
			node.Right = node.Right.Right
		} else {
			node.Right = nil
		}
		return true
	} else if node.Key < target {
		return delete(node.Right, target)
	} else if node.Key > target {
		return delete(node.Left, target)
	}

	return false
}

func (r *Root) update(target int, value int) bool {
	return update(r.Root, target, value)
}
func update(node *Node, target int, value int) bool {
	if node == nil {
		return false
	}
	if node.Key == target {
		node.Value = value
		return true
	} else if node.Key > target {
		return update(node.Left, target, value)
	} else {
		return update(node.Right, target, value)
	}

}
func (r *Root) DLR() {
	DLR(r.Root, 0)
}
func DLR(root *Node, depth int) {
	if root == nil {
		return
	}
	var block string
	for i := 0; i < depth; i++ {
		block += "\t\t"
	}

	fmt.Println(block + " -------------")
	fmt.Println(block+"|", root.Key)
	fmt.Println(block+"|", root.Value)
	fmt.Println(block + " -------------")
	DLR(root.Left, depth+1)
	DLR(root.Right, depth+1)
}

func (r *Root) LDR() {
	LDR(r.Root, 0)
}
func LDR(root *Node, depth int) {
	if root == nil {
		return
	}
	var block string
	for i := 0; i < depth; i++ {
		block += "\t\t"
	}
	LDR(root.Left, depth+1)
	fmt.Println(block + " -------------")
	fmt.Println(block+"|", root.Key)
	fmt.Println(block+"|", root.Value)
	fmt.Println(block + " -------------")
	LDR(root.Right, depth+1)
}

func (r *Root) LRD() {
	LRD(r.Root, 0)
}
func LRD(root *Node, depth int) {
	if root == nil {
		return
	}
	time.Sleep(time.Second)
	var block string
	for i := 0; i < depth; i++ {
		block += "\t\t"
	}
	LRD(root.Left, depth+1)
	LRD(root.Right, depth+1)
	fmt.Println(block + " -------------")
	fmt.Println(block+"|", root.Key)
	fmt.Println(block+"|", root.Value)
	fmt.Println(block + " -------------")

}

func (r *Root) LevelOrder() {
	queue := NewQueue()
	queue.Push(r.Root)
	var count int
	for queue.Count != 0{
		count++
		node:=queue.Pop()
		fmt.Println(node.Key)
		if node.Left !=nil{
			queue.Add(node.Left)
		}
		if node.Right != nil{
			queue.Add(node.Right)
		}
	}
}
