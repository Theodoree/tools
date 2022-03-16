package _300_1399

/*
1339. 分裂二叉树的最大乘积
给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
*/

func maxProduct(root *TreeNode) int {
	var ans float64
	dfs(root, sum(root), &ans)
	return int(ans) % (1e9 + 7)
}

func sum(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return sum(root.Left) + sum(root.Right) + float64(root.Val)
}

func dfs(root *TreeNode, sum float64, ans *float64) int {
	if root == nil {
		return 0
	}

	ls := dfs(root.Left, sum, ans)
	rs := dfs(root.Right, sum, ans)

	k := float64(ls + rs + root.Val)
	if (sum-k)*k > *ans {
		*ans = (sum - k) * k
	}

	return ls + rs + root.Val
}
