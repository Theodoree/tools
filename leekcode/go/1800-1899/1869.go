package _800_1899

/*
1869. 哪种连续子字符串更长
给你一个二进制字符串 s 。如果字符串中由 1 组成的 最长 连续子字符串 严格长于 由 0 组成的 最长 连续子字符串，返回 true ；否则，返回 false 。

例如，s = "110100010" 中，由 1 组成的最长连续子字符串的长度是 2 ，由 0 组成的最长连续子字符串的长度是 3 。
注意，如果字符串中不存在 0 ，此时认为由 0 组成的最长连续子字符串的长度是 0 。字符串中不存在 1 的情况也适用此规则。
*/

func checkZeroOnes(s string) bool {

	var max0, max1 int
	var count_1, count_0 int

	for i := 0; i < len(s); i++ {
		switch s[i] - '0' {
		case 0:
			if count_1 > 0 {
				if max1 < count_1 {
					max1 = count_1
				}
				count_1 = 0
			}
			count_0++
		case 1:
			if count_0 > 0 {
				if max0 < count_0 {
					max0 = count_0
				}
				count_0 = 0
			}
			count_1++
		}
	}

	if max1 < count_1 {
		max1 = count_1
	}

	if max0 < count_0 {
		max0 = count_0
	}

	return max1 > max0
}
