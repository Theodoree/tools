package _go


/*
剑指 Offer II 014. 字符串中的变位词
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {

	var alphabet [28]int16
	for i := 0; i < len(s1); i++ {
		alphabet[s1[i]-'a']++
	}

	for i := 0; i < len(s1); i++ {
		alphabet[s2[i]-'a']--
	}

	n := (*[7]int64)(unsafe.Pointer(&alphabet))
	if n[0]|n[1]|n[2]|n[3]|n[4]|n[5]|n[6] == 0 {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		alphabet[s2[i-len(s1)]-'a']++
		alphabet[s2[i]-'a']--
		if n[0]|n[1]|n[2]|n[3]|n[4]|n[5]|n[6] == 0 {
			return true
		}

	}

	return false
}
