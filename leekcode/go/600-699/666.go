package _00_699

/*
666. 路径总和 IV
对于一棵深度小于 5 的树，可以用一组三位十进制整数来表示。对于每个整数：
百位上的数字表示这个节点的深度 d，1 <= d <= 4。
十位上的数字表示这个节点在当前层所在的位置 P， 1 <= p <= 8。位置编号与一棵满二叉树的位置编号相同。
个位上的数字表示这个节点的权值 v，0 <= v <= 9。
给定一个包含三位整数的 升序 数组 nums ，表示一棵深度小于 5 的二叉树，请你返回 从根到所有叶子结点的路径之和 。
保证 给定的数组表示一个有效的连接二叉树。
*/
func pathSum(nums []int) int {
	tree := make([]int, 16)
	for i := range tree {
		tree[i] = -1
	}
	for i := range nums {
		pos, dep := nums[i]/10%10-1, nums[i]/100%10-1
		tree[1<<dep-1+pos] = nums[i] % 10 // 2^2 是异或……
	}
	return helper(tree, 0, 0)
}

func helper(tree []int, root, sum int) int {
	if tree[root] == -1 {
		return 0
	}
	if root >= 7 || tree[2*root+1] == -1 && tree[2*root+2] == -1 {  // 最后一层或者叶节点直接返回
		return tree[root] + sum
	}
	return helper(tree, 2*root+1, sum+tree[root]) + helper(tree, 2*root+2, sum+tree[root])
}

