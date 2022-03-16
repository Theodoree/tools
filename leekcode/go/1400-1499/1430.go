package _400_1499


/*
1430. 判断给定的序列是否是二叉树从根到叶的路径
给定一个二叉树，我们称从根节点到任意叶节点的任意路径中的节点值所构成的序列为该二叉树的一个 “有效序列” 。检查一个给定的序列是否是给定二叉树的一个 “有效序列” 。
我们以整数数组 arr 的形式给出这个序列。从根节点到任意叶节点的任意路径中的节点值所构成的序列都是这个二叉树的 “有效序列” 。
*/

func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil {
		return false
	}
	if len(arr) == 1 {
		return root.Left == nil&& root.Right == nil
	}
	var str string
	for _,v:=range arr{
		str += " " + strconv.Itoa(v)
	}

	return dfs(root,"",str)
}

func dfs(root *TreeNode,str string,target string) bool{
	if root == nil {
		return false
	}

	str += " " + strconv.Itoa(root.Val)
	if root .Left == nil && root.Right == nil && str == target{
		return true
	}
	return dfs(root.Left,str,target) || dfs(root.Right,str,target)
}

