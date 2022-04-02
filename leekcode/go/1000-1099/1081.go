package _000_1099



/*
1081. 不同字符的最小子序列
返回 s 字典序最小的子序列，该子序列包含 s 的所有不同字符，且只包含一次。
*/
func removeDuplicateLetters(s string) string {
	var buf [26]int
	for i := 0; i < len(s); i++ {
		buf[s[i]-'a']++
	}

	var flag int

	var h []byte

	for i := 0; i < len(s); i++ {
		buf[s[i]-'a']--
		if 1<<(s[i]-'a')&flag > 0 {
			continue
		}

		b := s[i] - 'a'

		for len(h) > 0 && h[len(h)-1] > b && buf[h[len(h)-1]] > 0 {
			flag = flag ^ 1<<h[len(h)-1]
			h = h[:len(h)-1]
		}

		flag |= 1 << b
		h = append(h, b)
	}

	var result []byte
	for i := 0; i < len(h); i++ {
		result = append(result, h[i]+'a')
	}

	return string(result)
}
