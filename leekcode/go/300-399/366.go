package _00_399

import (
	"fmt"
)

/*
366. 寻找完全二叉树的叶子节点

给你一棵完全二叉树，请按以下要求的顺序收集它的全部节点：

依次从左到右，每次收集并删除所有的叶子节点
重复如上过程直到整棵树为空
示例:

输入: [1,2,3,4,5]

          1
         / \
        2   3
       / \
      4   5

输出: [[4,5,3],[2],[1]]


解释:

1. 删除叶子节点 [4,5,3] ，得到如下树结构：

          1
         /
        2


2. 现在删去叶子节点 [2] ，得到如下树结构：

          1


3. 现在删去叶子节点 [1] ，得到空树：

          []
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findLeaves(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var val [][]int

	for {
		leaves := []int{}
		ok := LevelSort(root, root, &leaves)
		val = append(val, leaves)
		if ok {
			break
		}
	}

	return val
}

func LevelSort(root, parent *TreeNode, leaves *[]int) bool {

	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
		if root == parent.Left {
			parent.Left = nil
		} else if root == parent.Right {
			parent.Right = nil
		} else if root == parent {
			return true
		}
		return false
	}

	LevelSort(root.Left, root, leaves)
	LevelSort(root.Right, root, leaves)

	return false
}

func CreateTree(b []int) (*TreeNode, []*TreeNode) {

	var root *TreeNode
	root = &TreeNode{Val: b[0]}

	v := make([]*TreeNode, len(b)+1)
	var index, index1 int
	index = 1
	index1 = 1
	current := root
	v[0] = root
	for i := 1; i <= len(b)/2; i++ {

		if b[i*2-1] != 0 {
			current.Left = &TreeNode{Val: b[i*2-1]}
			v[index1] = current.Left
			index1++
		}
		if b[i*2] != 0 {
			current.Right = &TreeNode{Val: b[i*2]}
			v[index1] = current.Right
			index1++
		}

		current = v[index]
		index++
	}

	return root, v

}

func DLR(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf(" %d ", root.Val)
	DLR(root.Left)
	DLR(root.Right)
}

func main() {

	b := []int{37,
		-34, -48,
		0, -100, -100, 48,
		0, 0, 0, 0, -54, 0,
		-71, -22,
		0, 0, 0, 8}
	root, _ := CreateTree(b)

	// 0 1 2
	// 1 3 4
	// 2 5 6
	// 3 7 8
	// 4 9 10
	// 5 11 12
	// 6 13 14
	// 7 15 16
	// 8 17 18
	// fmt.Println(v[17].Val)
	// DLR(root)
	fmt.Println(findLeaves(root))

}
