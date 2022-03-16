package _400_1499

type Node struct {
	Val      int
	Children []*Node
	Left     *Node
	Right    *Node
	Next     *Node
}

/*
1490. 克隆 N 叉树
给定一棵 N 叉树的根节点 root ，返回该树的深拷贝（克隆）。
N 叉树的每个节点都包含一个值（ int ）和子节点的列表（ List[Node] ）。
*/
func cloneTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	n := &Node{Val: root.Val}

	var arr = []*Node{n, root}

	for len(arr) > 0 {
		src := arr[len(arr)-2]
		dist := arr[len(arr)-1]
		arr = arr[:len(arr)-2]
		for _, curDist := range dist.Children {
			curSrc := &Node{Val: curDist.Val}
			src.Children = append(src.Children, curSrc)
			arr = append(arr, curSrc, curDist)
		}
	}

	return n
}
