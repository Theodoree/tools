package _go

/*
剑指 Offer II 005. 单词长度的最大乘积
给定一个字符串数组 words，请计算当两个字符串 words[i] 和 words[j] 不包含相同字符时，它们长度的乘积的最大值。假设字符串中只包含英语的小写字母。如果没有不包含相同字符的一对字符串，返回 0。
*/
func maxProduct(words []string) int {
	n := len(words)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		len := len(words[i])
		for j := 0; j < len; j++ {
			arr[i] |= 1 << (words[i][j] - 'a')
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]&arr[j] == 0 {
				k := len(words[i]) * len(words[j])
				if ans < k {
					ans = k
				} else {
					k = ans
				}
			}
		}
	}
	return ans
}
