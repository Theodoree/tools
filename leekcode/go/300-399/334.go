package _00_399

/*
334. 递增的三元子序列
给你一个整数数组 nums ，判断这个数组中是否存在长度为 3 的递增子序列。
如果存在这样的三元组下标 (i, j, k) 且满足 i < j < k ，使得 nums[i] < nums[j] < nums[k] ，返回 true ；否则，返回 false 。
*/
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	var fir, sec int
	fir = math.MaxInt64
	sec = math.MaxInt64

	for i := 0; i < len(nums); i++ {
		if fir > nums[i] {
			fir = nums[i]
		} else if nums[i] > fir {
			if sec != math.MaxInt64 && nums[i] > sec {
				return true
			}
			if sec > nums[i] {
				sec = nums[i]
			}

		}
	}
	return false
}
