package _00_499

/*
459. 重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"

输出: False
示例 3:

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
通过次数82,459提交次数161,574
请问您在哪类招聘中遇到此题？
*/

func repeatedSubstringPattern(s string) bool {

	var subString string

	for i := 1; i < len(s); i++ {
		subString = s[:i]
		j := i
		for ; j < len(s); j += len(subString) {
			if j+len(subString) > len(s) {
				return false
			}
			if subString != s[j:j+len(subString)] {
				break
			}

		}
		if j == len(s) {
			return true
		}

	}

	return false
}
