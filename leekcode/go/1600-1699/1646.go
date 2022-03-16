package _600_1699

/*
1646. 获取生成数组中的最大值
给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：

nums[0] = 0
nums[1] = 1
当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]
返回生成数组 nums 中的 最大 值。
*/
func getMaximumGenerated(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 1
	}
	var buf = make([]int, n+1)
	buf[0] = 0
	buf[1] = 1

	var max int
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			buf[i] = buf[i/2]
		} else {
			buf[i] = buf[i/2] + buf[i/2+1]
		}
		if buf[i] > max {
			max = buf[i]
		}
	}
	return max
}
