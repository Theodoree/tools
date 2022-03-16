package _go

import "strings"

/*
剑指 Offer II 018. 有效的回文
给定一个字符串 s ，验证 s 是否是 回文串 ，只考虑字母和数字字符，可以忽略字母的大小写。
本题中，将空字符串定义为有效的 回文串 。
*/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		if !('a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9') {
			i++
			continue
		}
		if !('a' <= s[j] && s[j] <= 'z' || '0' <= s[j] && s[j] <= '9') {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
