package _900_1999

/*
1961. 检查字符串是否为数组前缀
给你一个字符串 s 和一个字符串数组 words ，请你判断 s 是否为 words 的 前缀字符串 。

字符串 s 要成为 words 的 前缀字符串 ，需要满足：s 可以由 words 中的前 k（k 为 正数 ）个字符串按顺序相连得到，且 k 不超过 words.length 。

如果 s 是 words 的 前缀字符串 ，返回 true ；否则，返回 false 。
*/
func isPrefixString(s string, words []string) bool {
	var result string
	for i := 0; i < len(words); i++ {
		result += words[i]
		if len(result) >= len(s) {
			return result == s
		}
	}
	return result == s
}
