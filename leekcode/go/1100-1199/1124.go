package _100_1199

/*
1124. 表现良好的最长时间段
给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
请你返回「表现良好时间段」的最大长度。
*/
func longestWPI(hours []int) int {
	for i, w := range hours {
		if w > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}

	prefix := make([]int, len(hours)+1)
	for i, w := range hours {
		prefix[i+1] = prefix[i] + w
	}

	stack := make([]int, 0)
	for i, w := range prefix {
		if len(stack) == 0 || prefix[stack[len(stack)-1]] > w {
			stack = append(stack, i)
		}
	}

	wpi := 0
	for i := len(prefix) - 1; i > wpi; i-- {
		for len(stack) > 0 && prefix[i] > prefix[stack[len(stack)-1]] {
			w := i - stack[len(stack)-1]
			if wpi < w {
				wpi = w
			}
			stack = stack[:len(stack)-1]
		}
	}
	return wpi
}
