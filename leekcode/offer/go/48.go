package _go

/*
剑指 Offer 48. 最长不含重复字符的子字符串
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
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
