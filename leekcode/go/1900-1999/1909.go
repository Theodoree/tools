package _900_1999

/*
1909. 删除一个元素使数组严格递增
给你一个下标从 0 开始的整数数组 nums ，如果 恰好 删除 一个 元素后，数组 严格递增 ，那么请你返回 true ，否则返回 false 。如果数组本身已经是严格递增的，请你也返回 true 。

数组 nums 是 严格递增 的定义为：对于任意下标的 1 <= i < nums.length 都满足 nums[i - 1] < nums[i] 。
*/
func canBeIncreasing(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var i int
	for i = 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			break
		}
	}

	if i == len(nums) {
		return true
	}

	if i == len(nums)-1 {
		return true
	}

	var j = i + 1
	for ; j < len(nums); j++ {

		if j == i+1 {
			if nums[j] > nums[j-2] {
				continue
			} else if nums[j] > nums[j-1] {
				if len(nums) == 3 {
					return true
				}

				if j >= 3 && nums[j-1] <= nums[j-3] {
					return false
				}
			} else {
				return false
			}
		} else if nums[j] <= nums[j-1] {
			return false
		}
	}

	return true
}
