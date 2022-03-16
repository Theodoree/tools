package _go



/*
剑指 Offer II 020. 回文子字符串的个数
给定一个字符串 s ，请计算这个字符串中有多少个回文子字符串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
*/
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2 * n - 1; i++ {
		l, r := i / 2, i / 2 + i % 2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}