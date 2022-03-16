package _700_1799

/*
1784. 检查二进制字符串字段
给你一个二进制字符串 s ，该字符串 不含前导零 。

如果 s 包含 零个或一个由连续的 '1' 组成的字段 ，返回 true​​​ 。否则，返回 false 。
*/
func checkOnesSegment(s string) bool {
	/*
		条件1: 零个1
		条件2: 或者全部连续连续的1
	*/
	var count int
	var flag bool
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			if flag {
				return false
			}
			count++
		} else {
			if count > 0 {
				flag = true
			}
		}

	}

	return true
}
