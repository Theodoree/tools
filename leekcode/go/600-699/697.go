package _00_699

/*
697. 数组的度
给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。
*/
//看明白了就不难，总结一下就是要找出数组的众数，并且还有找出众数在数组中第一次出现和最后一次出现的位置，
//两个位置组成区间长度就是答案, 如果众数不止一个，那么要取区间长度最短那个

func findShortestSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var m = make(map[int]int)
	var max int
	for _, v := range nums {
		m[v]++
		if m[v] > max {
			max = m[v]
		}
	}
	if max < 2 {
		return max
	}

	var min int
	min = math.MaxInt64
	for k, v := range m {
		if v != max {
			continue
		}

		firstIdx := 0
		for idx, num := range nums {
			if num == k {
				firstIdx = idx
				break
			}
		}

		for i := len(nums) - 1; i > firstIdx; i-- {
			if nums[i] != k {
				continue
			}

			if i-firstIdx+1 < min {
				min = i - firstIdx + 1
			}
			if min == 2 {
				return 2
			}

			break
		}

	}

	return min
}
