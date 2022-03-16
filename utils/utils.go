package utils

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ListNode) Create(val []int) *ListNode {
	if len(val) == 0 {
		return nil
	}

	head := &ListNode{Val: val[0]}
	cur := head

	for i := 1; i < len(val); i++ {
		cur.Next = &ListNode{Val: val[i]}
		cur = cur.Next
	}

	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (TreeNode) Create(b []int) (*TreeNode, []*TreeNode) {

	var root *TreeNode
	if len(b) == 0 {
		return nil, nil
	}
	root = &TreeNode{Val: b[0]}

	v := make([]*TreeNode, len(b)+1)
	var index, index1 int
	index = 1
	index1 = 1
	current := root
	v[0] = root
	for i := 1; i <= len(b)/2; i++ {

		if b[i*2-1] != Null {
			current.Left = &TreeNode{Val: b[i*2-1]}
			v[index1] = current.Left
			index1++
		}
		if i*2 < len(b) && b[i*2] != Null {
			current.Right = &TreeNode{Val: b[i*2]}
			v[index1] = current.Right
			index1++
		}
		current = v[index]
		index++
	}

	return root, v

}

func CreateTree(b []int) (*TreeNode, []*TreeNode) {
	return TreeNode{}.Create(b)
}
func CreateLinkNode(val []int) *ListNode {
	return ListNode{}.Create(val)
}

func PrintNode(n *ListNode) {

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}

}
func LevelTree(root *TreeNode, result *[][]int, depth int) {

	if root == nil {
		return
	}

	if len(*result) <= depth {
		*result = append(*result, []int{})
	}

	k := *result
	k[depth] = append(k[depth], root.Val)
	LevelTree(root.Left, result, depth+1)
	LevelTree(root.Right, result, depth+1)

}

func PrintTree(root *TreeNode, Type printType, array *[]int) {
	if root == nil {
		return
	}

	switch Type {
	case DLR:
		if array != nil {
			*array = append(*array, root.Val)
		} else {
			fmt.Printf(" %d ", root.Val)
		}
		PrintTree(root.Left, Type, array)
		PrintTree(root.Right, Type, array)
	case LDR:
		PrintTree(root.Left, Type, array)
		if array != nil {
			*array = append(*array, root.Val)
		} else {
			fmt.Printf(" %d ", root.Val)
		}
		PrintTree(root.Right, Type, array)
	case LRD:
		PrintTree(root.Left, Type, array)
		PrintTree(root.Right, Type, array)
		if array != nil {
			*array = append(*array, root.Val)
		} else {
			fmt.Printf(" %d ", root.Val)
		}
	case RDL:
		PrintTree(root.Right, Type, array)
		if array != nil {
			*array = append(*array, root.Val)
		} else {
			fmt.Printf(" %d ", root.Val)
		}
		PrintTree(root.Left, Type, array)
	case LEVEL:
		var v [][]int
		LevelTree(root, &v, 0)
		for _, val := range v {
			fmt.Println(val)
		}
	}
}

type printType string

const (
	DLR   printType = "DLR"
	LDR   printType = "LDR"
	LRD   printType = "LRD"
	RDL   printType = "RDL"
	LEVEL printType = "LEVEL"
	Null            = 0x7777777
)
