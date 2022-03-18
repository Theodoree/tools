package _00_599


/*
567. 字符串的排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
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

	// int64 = int16*4       4*
	for i := len(s1); i < len(s2); i++ {
		alphabet[s2[i-len(s1)]-'a']++
		alphabet[s2[i]-'a']--
		if n[0]|n[1]|n[2]|n[3]|n[4]|n[5]|n[6] == 0 {
			return true
		}

	}

	return false
}

