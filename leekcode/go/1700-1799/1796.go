package _700_1799

/*
1796. 字符串中第二大的数字
给你一个混合字符串 s ，请你返回 s 中 第二大 的数字，如果不存在第二大的数字，请你返回 -1 。
混合字符串 由小写英文字母和数字组成。
*/
func secondHighest(s string) int {
	var buf [10]int

	for _, v := range s {
		if v >= '0' && v <= '9' {
			buf[v-'0']++
		}
	}

	var flag bool
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == 0 {
			continue
		}
		if flag {
			return i
		}
		flag = true
	}

	return -1
}
