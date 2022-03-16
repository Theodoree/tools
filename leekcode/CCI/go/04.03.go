package _go

/*
面试题 04.03. 特定深度节点链表
给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。



示例：

输入：[1,2,3,4,5,null,7,8]

        1
       /  \
      2    3
     / \    \
    4   5    7
   /
  8

输出：[[1],[2,3],[4,5,7],[8]]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	var result []*ListNode
	var tail []*ListNode
	dfs(0, tree, &result, &tail)
	return result
}

func dfs(depth int, root *TreeNode, result, tail *[]*ListNode) {
	if root == nil {
		return
	}

	l := &ListNode{
		Val:  root.Val,
		Next: nil,
	}
	if len(*result) <= depth {
		*result = append(*result, l)
		*tail = append(*tail, l)
	} else {
		(*tail)[depth].Next = l
		(*tail)[depth] = l
	}
	dfs(depth+1, root.Left, result, tail)
	dfs(depth+1, root.Right, result, tail)
}

