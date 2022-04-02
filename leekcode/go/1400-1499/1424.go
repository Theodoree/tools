package _400_1499


/*
1424. 对角线遍历 II
给你一个列表 nums ，里面每一个元素都是一个整数列表。请你依照下面各图的规则，按顺序返回 nums 中对角线上的整数。
*/

func findDiagonalOrder(nums [][]int) []int {
	var ans []int
	m, index := len(nums), -1
	for i := 0; i < m; i++ {
		next := i
		ans = append(ans, nums[i][0])
		nums[next] = nums[i][1:]
		if len(nums[next]) != 0 {
			next--
		}
		for j := i-1; j > index; j-- {
			ans = append(ans, nums[j][0])
			nums[next] = nums[j][1:]
			if len(nums[next]) != 0 {
				next--
			}
		}
		index = next
	}
	for index != m-1 {
		next := m-1
		for j := m-1; j > index; j-- {
			ans = append(ans, nums[j][0])
			nums[next] = nums[j][1:]
			if len(nums[next]) != 0 {
				next--
			}
		}
		index = next
	}
	return ans
}
