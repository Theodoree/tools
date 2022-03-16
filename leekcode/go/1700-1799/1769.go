package _700_1799


/*
1769. 移动所有球到每个盒子所需的最小操作数
有 n 个盒子。给你一个长度为 n 的二进制字符串 boxes ，其中 boxes[i] 的值为 '0' 表示第 i 个盒子是 空 的，而 boxes[i] 的值为 '1' 表示盒子里有 一个 小球。
在一步操作中，你可以将 一个 小球从某个盒子移动到一个与之相邻的盒子中。第 i 个盒子和第 j 个盒子相邻需满足 abs(i - j) == 1 。注意，操作执行后，某些盒子中可能会存在不止一个小球。
返回一个长度为 n 的数组 answer ，其中 answer[i] 是将所有小球移动到第 i 个盒子所需的 最小 操作数。
每个 answer[i] 都需要根据盒子的 初始状态 进行计算。
*/

func minOperations(boxes string) []int {
	n := len(boxes)
	ans := make([]int, n)
	leftCount := 0
	rightCount := 0
	for i, c := range boxes {
		if c == '1' {
			rightCount++
			ans[0] += i
		}
	}

	for i := 1; i < n; i++ {
		if boxes[i - 1] == '1' {
			leftCount++
			rightCount--
		}
		ans[i] = ans[i - 1] + leftCount - rightCount
	}
	return ans
}
