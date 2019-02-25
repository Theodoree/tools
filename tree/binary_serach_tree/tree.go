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
	Key         int64
	Value       int64
}

func NewRoot(key int64, value int64) *Root {
	return &Root{
		NewNode(key, value), 1,
	}

}

func NewNode(key int64, value int64) *Node {
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

func (r *Root) Insert(root *Node, key int64, value int64) *Node {
	if root == nil {
		r.Count++
		node := NewNode(key, value)
		return node
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

func (r *Root) Contain(key int64) bool {
	return contain(r.Root, key)
}
func contain(node *Node, key int64) bool {
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

func (r *Root) Search(root *Node, key int64) *Node {
	return search(root, key)
}
func search(root *Node, key int64) *Node {
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

func (r *Root) Delete(target int64) *Node {
	node := delete(r.Root, target)
	if node != nil {
		r.Count--
	}
	return node
}

//删除指定节点 返回子树root
func delete(node *Node, target int64) *Node {
	if node == nil {
		return nil
	}

	if node.Key == target {
		minNode,minNodeParent := min(node.Right)
		minNodeParent.Left = nil
		minNode.Left = node.Left
		minNode.Right = node.Right
		return minNode
	}

	if node.Key > target {
		node.Right = delete(node.Right, target)
		return node
	}

	if node.Key < target {
		node.Left = delete(node.Left, target)
		return node
	}

	return nil
}

func (r *Root) DeleteMin() *Node {
	if node := deleteMin(r.Root); node != nil {
		r.Count --
		return node
	}
	return nil
}
func deleteMin(node *Node) *Node {

	if node.Left == nil {
		right := node.Right
		return right
	}
	node.Left = deleteMin(node.Left)
	return node
}
func (r *Root) DeleteMax() *Node {
	if node := deleteMax(r.Root); node != nil {
		r.Count--
		return node
	}
	return nil
}
func deleteMax(node *Node) *Node {
	if node.Right == nil {
		left := node.Left
		return left
	}
	node.Right = deleteMax(node.Right)
	return node
}

func (r *Root) update(target int64, value int64) *Node {
	return update(r.Root, target, value)
}
func update(node *Node, target int64, value int64) *Node {
	if node == nil {
		return node
	}
	if node.Key == target {
		node.Value = value
		return node
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
	levelOrder(r.Root)
}

func levelOrder(node *Node) {
	queue := NewQueue()
	queue.Push(node)
	var count int
	for queue.Count != 0 {
		count++
		node := queue.Pop()
		fmt.Println(node.Key)
		if node.Left != nil {
			queue.Add(node.Left)
		}
		if node.Right != nil {
			queue.Add(node.Right)
		}
	}

}
func (r *Root) Max() *Node {
	return max(r.Root)
}

func max(node *Node) *Node {
	current := node
	for {
		if current.Right != nil {
			current = current.Right
		} else {
			return current
		}
	}

}

func (r *Root) Min() (*Node, *Node) {
	return min(r.Root)
}

func min(node *Node) (*Node, *Node) {
	current := node
	var parten *Node
	for {
		if current.Left != nil {
			parten = current
			current = current.Left
		} else {
			return current, parten
		}
	}
}
