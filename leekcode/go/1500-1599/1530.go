package _500_1599



/*
1530. 好叶子节点对的数量
给你二叉树的根节点 root 和一个整数 distance 。
如果二叉树中两个 叶 节点之间的 最短路径长度 小于或者等于 distance ，那它们就可以构成一组 好叶子节点对 。
返回树中 好叶子节点对的数量 。
*/
func countPairs(root *TreeNode, distance int) (res int) {
	// 后序, 处理返回来的左右到叶子的路径数组, 返回来一次, 数组里的数都要+1
	// 遍历查找小于distance的距离和即可
	var dfs func(root *TreeNode) []int
	dfs = func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		if root.Left == nil && root.Right == nil {
			return []int{0}
		}
		lPath := dfs(root.Left)
		for i := range lPath {
			lPath[i]++
		}
		rPath := dfs(root.Right)
		for i := range rPath {
			rPath[i]++
		}
		for _, l := range lPath {
			for _, r := range rPath {
				if l+r <= distance {
					res++
				}
			}
		}
		return append(lPath, rPath...)
	}
	dfs(root)
	return
}
