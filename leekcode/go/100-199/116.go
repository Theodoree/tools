package _00_199


/*
116. 填充每个节点的下一个右侧节点指针
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

*/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var fn func(r *Node, result *[][]*Node, depth int)

	fn = func(r *Node, result *[][]*Node, depth int) {
		if r == nil {
			return
		}

		if len(*result) <= depth {
			*result = append(*result, []*Node{})
		}
		(*result)[depth] = append((*result)[depth], r)

		fn(r.Left, result, depth+1)
		fn(r.Right, result, depth+1)
	}

	root.Next = nil
	var result [][]*Node
	fn(root, &result, 0)

	for i := 1; i < len(result); i++ {
		for j := 0; j < len(result[i])-1; j++ {
			result[i][j].Next = result[i][j+1]
		}
	}

	return root

}
