package _00_599



/*
510. 二叉搜索树中的中序后继 II
给定一棵二叉搜索树和其中的一个节点 node ，找到该节点在树中的中序后继。如果节点没有中序后继，请返回 null 。
一个节点 node 的中序后继是键值比 node.val 大所有的节点中键值最小的那个。
你可以直接访问结点，但无法直接访问树。每个节点都会有其父节点的引用。节点 Node 定义如下：
*/

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right != nil { // 有右子树
		res := node.Right
		for res.Left != nil {
			res = res.Left
		}
		return res
	} else if node.Parent != nil {
		for node.Parent != nil {
			if node.Parent.Val > node.Val {
				return node.Parent
			}
			node = node.Parent
		}
	}
	return nil
}
