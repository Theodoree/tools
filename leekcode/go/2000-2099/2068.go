package _000_2099

/*
2068. 检查两个字符串是否几乎相等
如果两个字符串 word1 和 word2 中从 'a' 到 'z' 每一个字母出现频率之差都 不超过 3 ，那么我们称这两个字符串 word1 和 word2 几乎相等 。
给你两个长度都为 n 的字符串 word1 和 word2 ，如果 word1 和 word2 几乎相等 ，请你返回 true ，否则返回 false 。
一个字母 x 的出现 频率 指的是它在字符串中出现的次数。
*/
func checkAlmostEquivalent(word1 string, word2 string) bool {
	var buf [26]int

	for _, v := range word1 {
		buf[v-'a']++
	}

	for _, v := range word2 {
		buf[v-'a']--
	}

	for _, v := range buf {
		if v > 3 || v < -3 {
			return false
		}
	}
	return true
}
