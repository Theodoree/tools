package _00_299



/*
272. 最接近的二叉搜索树值 II
给定二叉搜索树的根 root 、一个目标值 target 和一个整数 k ，返回BST中最接近目标的 k 个值。你可以按 任意顺序 返回答案。
题目 保证 该二叉搜索树中只会存在一种 k 个值集合最接近 target
*/
func closestKValues(root *TreeNode, target float64, k int) []int {

	if root == nil || k == 0 {
		return nil
	}

	var result []int
	GetSlice(root, &result)
	idx := 0

	if k == 1 {
		min := float64(math.MaxInt64)
		for i := 0; i < len(result); i++ {
			if math.Abs(target-float64(result[i])) < min {
				min = math.Abs(target - float64(result[i]))
				idx = i
			}
		}
		return []int{result[idx]}
	}

	for i := 0; i < len(result)-k; i++ {
		if math.Abs(target-float64(result[idx])) > math.Abs(target-float64(result[i+k])) {
			idx = i+1
		}

	}

	return result[idx : idx+k]
}

func GetSlice(root *TreeNode, f *[]int) {

	if root == nil {
		return
	}

	GetSlice(root.Left, f)
	*f = append(*f, root.Val)
	GetSlice(root.Right, f)

}
