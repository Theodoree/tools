package _800_1899

/*
1848. 到目标元素的最小距离
给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数 target 和 start ，请你找出一个下标 i ，满足 nums[i] == target 且 abs(i - start) 最小化 。注意：abs(x) 表示 x 的绝对值。
*/

func getMinDistance(nums []int, target int, start int) int {

	if nums[start] == target {
		return 0
	}

	// 先看左边
	left := start - 1
	for left >= 0 {
		if nums[left] == target {
			break
		}
		left--
	}

	right := start + 1
	for right < len(nums) {
		if nums[right] == target {
			break
		}
		right++
	}

	n1 := int(math.Abs(float64(left - start)))
	n2 := int(math.Abs(float64(right - start)))

	if left == -1 {
		return n2
	}
	if right == len(nums) {
		return n1
	}

	if n1 < n2 {
		return n1
	}

	return n2
}
