package _200_1299

/*
1261. 在受污染的二叉树中查找元素
给出一个满足下述规则的二叉树：

root.val == 0
如果 treeNode.val == x 且 treeNode.left != null，那么 treeNode.left.val == 2 * x + 1
如果 treeNode.val == x 且 treeNode.right != null，那么 treeNode.right.val == 2 * x + 2
现在这个二叉树受到「污染」，所有的 treeNode.val 都变成了 -1。

请你先还原二叉树，然后实现 FindElements 类：
FindElements(TreeNode* root) 用受污染的二叉树初始化对象，你需要先把它还原。
bool find(int target) 判断目标值 target 是否存在于还原后的二叉树中并返回结果。
*/


type FindElements struct {
	m map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	var f FindElements
	root.Val = 0
	f.m=make(map[int]struct{})
	recovery(root,f.m)
	return f
}

func recovery(root *TreeNode,m map[int]struct {}){
	if root == nil {
		return
	}
	m[root.Val] = struct {}{}

	if root.Left != nil {
		root.Left.Val = 2*root.Val+1
	}
	if root.Right != nil {
		root.Right.Val = 2*root.Val+2
	}

	recovery(root.Left,m)
	recovery(root.Right,m)


}

func (f *FindElements) Find(target int) bool {
	_,ok:=f.m[target]
	return ok
}

