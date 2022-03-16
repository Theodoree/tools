package _400_1499

/*
1470. 重新排列数组
给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。
请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。
*/
func shuffle(nums []int, n int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i < n {
			result[i*2] = nums[i]
		} else {
			result[(i-n)*2+1] = nums[i]
		}

	}
	return result
}
