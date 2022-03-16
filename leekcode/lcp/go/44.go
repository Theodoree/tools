package _go

/*
LCP 44. 开幕式焰火
「力扣挑战赛」开幕式开始了，空中绽放了一颗二叉树形的巨型焰火。
给定一棵二叉树 root 代表焰火，节点值表示巨型焰火这一位置的颜色种类。请帮小扣计算巨型焰火有多少种不同的颜色。
*/
func numColor(root *TreeNode) int {
	var m = make(map[int]struct{})
	bfs(root, m)
	return len(m)
}

func bfs(root *TreeNode, m map[int]struct{}) {
	if root == nil {
		return
	}
	m[root.Val] = struct{}{}
	bfs(root.Left, m)
	bfs(root.Right, m)

}
