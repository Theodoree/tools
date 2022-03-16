package _500_1599

type Node struct {
	Val      int
	Children []*Node
	Left     *Node
	Right    *Node
	Next     *Node
}

/*
1506. 找到 N 叉树的根节点
给定一棵 N 叉树 的所有节点在一个数组  Node[] tree 中，树中每个节点都有 唯一的值 。
找到并返回 N 叉树的 根节点 。
*/
func findRoot(tree []*Node) *Node {
	if len(tree) == 0 {
		return nil
	}
	var m = make(map[int]*Node)
	// child = parent
	for _, v := range tree {
		for _, val := range v.Children {
			m[val.Val] = v
		}
	}

	n := tree[0]
	for m[n.Val] != nil {
		n = m[n.Val]
	}
	return n
}
