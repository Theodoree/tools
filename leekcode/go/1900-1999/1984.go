package _900_1999

/*
1984. 学生分数的最小差值
给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。另给你一个整数 k 。
从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。
返回可能的 最小差值 。
*/
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	less := math.MaxInt64
	for i := 0; i <= len(nums)-k; i++ {
		var sum int
		for j := i + 1; j < i+k; j++ {
			sum += nums[j] - nums[j-1]
		}

		if sum < less {
			less = sum
		}

	}

	if less == math.MaxInt64 {
		return 0
	}
	return less
}

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ret := 100_000_00
	for i := 0; i+k-1 < len(nums); i++ {
		ret = min(ret, nums[i+k-1]-nums[i])
	}
	return ret
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
