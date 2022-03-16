package _500_1599

/*
  1522. N 叉树的直径
  给定一棵 N 叉树的根节点 root ，计算这棵树的直径长度。

  N 叉树的直径指的是树中任意两个节点间路径中 最长 路径的长度。这条路径可能经过根节点，也可能不经过根节点。

  （N 叉树的输入序列以层序遍历的形式给出，每组子节点用 null 分隔）


*/
func diameter(root *Node) int {

	var max int
	dfs(root, &max)
	return max
}
func dfs(root *Node, max *int) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}

	var arr []int
	for _, v := range root.Children {
		arr = append(arr, dfs(v, max))
	}
	sort.Ints(arr)

	if len(arr) >= 2 {
		if arr[len(arr)-1]+arr[len(arr)-2] > *max {
			*max = arr[len(arr)-1] + arr[len(arr)-2]
		}
	}
	if arr[len(arr)-1] > *max {
		*max = arr[len(arr)-1]
	}

	return arr[len(arr)-1] + 1
}
