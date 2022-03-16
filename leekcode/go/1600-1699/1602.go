package _600_1699



/*
1602. 找到二叉树中最近的右侧节点
给定一棵二叉树的根节点 root 和树中的一个节点 u ，返回与 u 所在层中距离最近的右侧节点，当 u 是所在层中最右侧的节点，返回 null 。
*/
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == u {
		return nil
	}

	var uDepth int
	uDepth = -1
	var target *TreeNode
	levelOrder(root,u,0,&uDepth,&target)
	return target
}


func levelOrder(root *TreeNode,u *TreeNode,depth int,uDepth *int,t **TreeNode)  {
	if root == nil {
		return
	}
	if root == u {
		*uDepth = depth
		return
	}
	levelOrder(root.Left,u,depth+1,uDepth,t)
	if *uDepth == depth && *t == nil{
		*t = root
		return
	}

	levelOrder(root.Right,u,depth+1,uDepth,t)
}

