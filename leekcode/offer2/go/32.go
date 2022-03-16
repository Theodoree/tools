package _go

/*
剑指 Offer II 032. 有效的变位词
给定两个字符串 s 和 t ，编写一个函数来判断它们是不是一组变位词（字母异位词）。
注意：若 s 和 t 中每个字符出现的次数都相同且字符顺序不完全相同，则称 s 和 t 互为变位词（字母异位词）。
*/
func isAnagram(s string, t string) bool {
	if s == t || len(s) != len(t) {
		return false
	}
	var buf [26]int
	for _, v := range s {
		buf[v-'a']++
	}
	for _, v := range t {
		buf[v-'a']--
	}
	for _, v := range buf {
		if v < 0 {
			return false
		}
	}
	return true
}
