package _200_1299

/*
1239. 串联字符串的最大长度
给定一个字符串数组 arr，字符串 s 是将 arr 某一子序列字符串连接所得的字符串，如果 s 中的每一个字符都只出现过一次，那么它就是一个可行解。
请返回所有可行解 s 中最长长度。
*/
func maxLength(arr []string) int {
	// 可以改dp,没必要
	fn := func(s []string) []int {
		var result = make([]int, 0, len(s))
		for _, str := range s {
			var f int
			for _, v := range str {
				if f&(1<<(v-'a')) > 0 {
					f = -1
					break
				}
				f |= 1 << (v - 'a')
			}
			result = append(result, f)
		}
		return result
	}
	var max int
	backtracking(arr, fn(arr), 0, 0, 0, &max)

	return max
}

func backtracking(arr []string, nums []int, idx int, alphabet, count int, max *int) {
	if count > *max {
		*max = count
	}
	if idx == len(arr) {
		return
	}

	if nums[idx] == -1 || alphabet&nums[idx] > 0 {
		backtracking(arr, nums, idx+1, alphabet, count, max)
		return
	}
	backtracking(arr, nums, idx+1, alphabet|nums[idx], count+len(arr[idx]), max)
	backtracking(arr, nums, idx+1, alphabet, count, max)
}
