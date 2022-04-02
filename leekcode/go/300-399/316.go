package _00_399



/*
316. 去除重复字母
给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
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

