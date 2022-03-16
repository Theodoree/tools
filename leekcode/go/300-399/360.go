package _00_399



/*
360. 有序转化数组
给你一个已经 排好序 的整数数组 nums 和整数 a 、 b 、 c 。对于数组中的每一个元素 nums[i] ，计算函数值 f(x) = ax2 + bx + c ，请 按升序返回数组 。
*/
func sortTransformedArray(nums []int, a int, b int, c int) []int {
	l := len(nums)
	var result []int
	left, right := 0, l-1
	memo := make(map[int]int)
	// 提前算好每个位置数字的运算式结果
	for k, v := range nums {
		memo[k] = a*v*v + b*v + c
	}
	for left <= right {
		leftValue, rightValue := memo[left], memo[right]
		if a > 0 { // 函数存在最小值，越靠边上，越大
			if leftValue < rightValue {
				result = append([]int{rightValue}, result...)
				right--
			} else {
				result = append([]int{leftValue}, result...)
				left++
			}
		} else { // 函数存在最大值，越靠边上，越小
			if leftValue < rightValue {
				result = append(result, leftValue)
				left++
			} else {
				result = append(result, rightValue)
				right--
			}
		}
	}
	return result
}

