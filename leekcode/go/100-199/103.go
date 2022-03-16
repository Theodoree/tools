package _00_199


/*
103. 二叉树的锯齿形层序遍历
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

func zigzagLevelOrder(root *TreeNode) [][]int {

	var result [][]int
	bfs(root, &result, 0)
	return result
}

func bfs(r *TreeNode, result *[][]int, depth int) {
	if r == nil {
		return
	}

	if len(*result) <= depth {
		*result = append(*result, []int{})
	}

	if depth %2 == 0 {
		(*result)[depth] = append((*result)[depth], r.Val)
	}else{
		(*result)[depth] = append([]int{r.Val},(*result)[depth]...)
	}

	depth++
	bfs(r.Left, result, depth)
	bfs(r.Right, result, depth)

}

