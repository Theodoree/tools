package _900_1999

/*
1941. 检查是否所有字符出现次数相同
给你一个字符串 s ，如果 s 是一个 好 字符串，请你返回 true ，否则请返回 false 。
如果 s 中出现过的 所有 字符的出现次数 相同 ，那么我们称字符串 s 是 好 字符串。
*/
func areOccurrencesEqual(s string) bool {
	var buf [26]int
	for _, v := range s {
		buf[v-'a']++
	}
	var cnt int
	for _, v := range buf {
		if v == 0 {
			continue
		}
		if cnt == 0 {
			cnt = v
		}

		if v != cnt {
			return false
		}

	}
	return true

}
