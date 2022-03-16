package _600_1699

/*
1624. 两个相同字符之间的最长子字符串
给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。
子字符串 是字符串中的一个连续字符序列。
*/
func maxLengthBetweenEqualCharacters(s string) int {

	if len(s) == 2 && s[0] == s[1] {
		return 0
	}

	var count int
	for i := 0; i < len(s); i++ {

		for j := len(s) - 1; j > i; j-- {
			if s[i] == s[j] {
				if j-i > count {
					count = j - i - 1
				}
				break
			}
		}

		if len(s)-i <= count {
			return count
		}
	}

	return -1
}
