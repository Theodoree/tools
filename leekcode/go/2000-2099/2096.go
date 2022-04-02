package _000_2099



/*
2096. 从二叉树一个节点到另一个节点每一步的方向
给你一棵 二叉树 的根节点 root ，这棵二叉树总共有 n 个节点。每个节点的值为 1 到 n 中的一个整数，且互不相同。给你一个整数 startValue ，表示起点节点 s 的值，和另一个不同的整数 destValue ，表示终点节点 t 的值。

请找到从节点 s 到节点 t 的 最短路径 ，并以字符串的形式返回每一步的方向。每一步用 大写 字母 'L' ，'R' 和 'U' 分别表示一种方向：

'L' 表示从一个节点前往它的 左孩子 节点。
'R' 表示从一个节点前往它的 右孩子 节点。
'U' 表示从一个节点前往它的 父 节点。
请你返回从 s 到 t 最短路径 每一步的方向。
*/
func getDirections(root *TreeNode, startValue, destValue int) string {
	var path []byte
	var target int
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node.Val == target {
			return true
		}
		if node.Left != nil {
			path = append(path, 'L')
			if dfs(node.Left) {
				return true
			}
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			path = append(path, 'R')
			if dfs(node.Right) {
				return true
			}
			path = path[:len(path)-1]
		}
		return false
	}
	target = startValue
	dfs(root)
	pathToStart := path

	path = nil
	target = destValue
	dfs(root)
	pathToDest := path

	for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
		pathToStart = pathToStart[1:]
		pathToDest = pathToDest[1:]
	}

	for i := range pathToStart {
		pathToStart[i] = 'U'
	}
	pathToStart = append(pathToStart, pathToDest...)
	return *(*string)(unsafe.Pointer(&pathToStart))
}
func init() { debug.SetGCPercent(-1) }

