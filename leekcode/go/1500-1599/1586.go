package _500_1599


/*
1586. 二叉搜索树迭代器 II
实现二叉搜索树（BST）的中序遍历迭代器 BSTIterator 类：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的实例。二叉搜索树的根节点 root 作为构造函数的参数传入。内部指针使用一个不存在于树中且小于树中任意值的数值来初始化。
boolean hasNext() 如果当前指针在中序遍历序列中，存在右侧数值，返回 true ，否则返回 false 。
int next() 将指针在中序遍历序列中向右移动，然后返回移动后指针所指数值。
boolean hasPrev() 如果当前指针在中序遍历序列中，存在左侧数值，返回 true ，否则返回 false 。
int prev() 将指针在中序遍历序列中向左移动，然后返回移动后指针所指数值。
注意，虽然我们使用树中不存在的最小值来初始化内部指针，第一次调用 next() 需要返回二叉搜索树中最小的元素。
你可以假设 next() 和 prev() 的调用总是有效的。即，当 next()/prev() 被调用的时候，在中序遍历序列中一定存在下一个/上一个元素。
*/


type BSTIterator struct {
	arr []int
	idx  int
}

func ldr(root *TreeNode,result *[]int){
	if root == nil {
		return
	}
	ldr(root.Left,result)
	*result = append(*result,root.Val)
	ldr(root.Right,result)
}

func Constructor(root *TreeNode) BSTIterator {
	var b BSTIterator
	ldr(root,&b.arr)
	b.idx = -1
	return b
}


func (b *BSTIterator) HasNext() bool {
	return b.idx < len(b.arr) -1
}


func (b *BSTIterator) Next() int {
	b.idx++
	val:=b.arr[b.idx]
	return val
}


func (b *BSTIterator) HasPrev() bool {
	return b.idx > 0
}


func (b *BSTIterator) Prev() int {
	b.idx--
	return b.arr[b.idx]
}

