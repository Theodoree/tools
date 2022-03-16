package _500_1599

/*
1544. 整理字符串
给你一个由大小写英文字母组成的字符串 s 。
一个整理好的字符串中，两个相邻字符 s[i] 和 s[i+1]，其中 0<= i <= s.length-2 ，要满足如下条件:
若 s[i] 是小写字符，则 s[i+1] 不可以是相同的大写字符。
若 s[i] 是大写字符，则 s[i+1] 不可以是相同的小写字符。
请你将字符串整理好，每次你都可以从字符串中选出满足上述条件的 两个相邻 字符并删除，直到字符串整理好为止。
请返回整理好的 字符串 。题目保证在给出的约束条件下，测试样例对应的答案是唯一的。
注意：空字符串也属于整理好的字符串，尽管其中没有任何字符。
*/
func makeGood(s string) string {
	buf := []byte(s)
	count := 0
	for count != len(buf) {
		flag := false

		for i := 0; i < len(buf)-1; i++ {
			if buf[i] == 0 {
				continue
			}
			nextIdx := i + 1
			for nextIdx < len(buf) && buf[nextIdx] == 0 {
				nextIdx++
			}
			if nextIdx == len(buf) {
				break
			}
			if (buf[i] >= 'a' && buf[nextIdx] == buf[i]-32) || (buf[i] >= 'A' && buf[nextIdx] == buf[i]+32) {
				buf[i] = 0
				buf[nextIdx] = 0
				count += 2
				flag = true
				break
			}

		}

		if !flag {
			break
		}
	}

	if count == len(buf) {
		return ""
	}
	result := make([]byte, 0, len(buf)-count)
	for _, v := range buf {
		if v == 0 {
			continue
		}
		result = append(result, v)
	}

	return BytesToString(result)
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
