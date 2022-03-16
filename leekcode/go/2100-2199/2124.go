package _100_2199

/*
2124. 检查是否所有 A 都在 B 之前
给你一个 仅 由字符 'a' 和 'b' 组成的字符串  s 。如果字符串中 每个 'a' 都出现在 每个 'b' 之前，返回 true ；否则，返回 false 。



示例 1：

输入：s = "aaabbb"
输出：true
解释：
'a' 位于下标 0、1 和 2 ；而 'b' 位于下标 3、4 和 5 。
因此，每个 'a' 都出现在每个 'b' 之前，所以返回 true 。
*/
func checkString(s string) bool {
	if len(s) == 0 {
		return true
	}
	var ok [2]bool // 是否有a 是否有b
	for _, v := range s {
		switch v {
		case 'a':
			if ok[1] {
				return false
			}
			ok[0] = true
		case 'b':
			ok[1] = true
		}
	}

	return ok[0] || ok[1]
}
