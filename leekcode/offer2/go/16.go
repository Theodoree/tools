package _go

/*
剑指 Offer II 016. 不含重复字符的最长子字符串
给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。
*/
func lengthOfLongestSubstring(s string) int {
	var lastOccurred [256]int
	for i := 0; i < len(lastOccurred); i++ {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastOccurred[ch] != -1 && lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
