package _00_499

/*
451. 根据字符出现频率排序
给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
*/
func frequencySort(s string) string {
	var buf [62]int
	for _, v := range s {
		if v >= 'a' {
			buf[v-'a']++
		} else if v >= 'A' {
			buf[v-'A'+26]++
		} else {
			buf[v-'0'+52]++
		}
	}

	var builder strings.Builder
	for {
		max := 0
		idx := -1
		for i := 0; i < len(buf); i++ {
			if buf[i] > max {
				max = buf[i]
				idx = i
			}
		}
		if idx == -1 {
			break
		}
		buf[idx] = -1

		for i := 0; i < max; i++ {
			if idx >= 52 {
				builder.WriteByte('0' + byte(idx-52))
			} else if idx >= 26 {
				builder.WriteByte('A' + byte(idx-26))
			} else {
				builder.WriteByte('a' + byte(idx))
			}
		}

	}

	return builder.String()
}
