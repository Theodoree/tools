package _700_1799

/*
1763. 最长的美好子字符串
当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。比方说，"abABB" 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。
给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。如果有多个答案，请你返回 最早 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。
*/
func longestNiceSubstring(s string) string {

	var subStr string
	for i := 0; i < len(s); i++ {
		var count, bigcount int
		for j := i; j < len(s); j++ {
			switch {
			case s[j] >= 'a':
				count |= 1 << (s[j] - 'a')
			case s[j] >= 'A':
				bigcount |= 1 << (s[j] - 'A')
			}

			if count > 0 && count == bigcount && j-i >= len(subStr) {
				subStr = s[i : j+1]
			}
		}
	}

	return subStr

}
