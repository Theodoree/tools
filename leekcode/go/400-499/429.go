package _00_499

/*
429. N 叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
*/
func levelOrder(root *Node) [][]int {
	var result [][]int
	level(root, 0, &result)
	return result
}

func level(root *Node, depth int, result *[][]int) {
	if root == nil {
		return
	}
	if depth <= len(*result) {
		*result = append(*result, []int{})
	}

	(*result)[depth] = append((*result)[depth], root.Val)

	for _, v := range root.Children {
		level(v, depth+1, result)
	}

}
