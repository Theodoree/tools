package _400_1499

/*
1427. 字符串的左右移
给定一个包含小写英文字母的字符串 s 以及一个矩阵 shift，其中 shift[i] = [direction, amount]：
direction 可以为 0 （表示左移）或 1 （表示右移）。
amount 表示 s 左右移的位数。
左移 1 位表示移除 s 的第一个字符，并将该字符插入到 s 的结尾。
类似地，右移 1 位表示移除 s 的最后一个字符，并将该字符插入到 s 的开头。
对这个字符串进行所有操作后，返回最终结果。
*/
func stringShift(s string, shift [][]int) string {
	if len(shift) == 0 {
		return s
	}

	var amount int

	for i := 0; i < len(shift); i++ {
		switch shift[i][0] {
		case 0:
			amount -= shift[i][1]
		case 1:
			amount += shift[i][1]
		}
	}
	if amount == 0 {
		return s
	}
	amount %= len(s)

	buf := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if amount < 0 { // 左移
			idx := (i + amount*(-1)) % len(s)
			buf = append(buf, s[idx])
		} else { // 右移
			idx := (len(s) - amount + i) % len(s)
			buf = append(buf, s[idx])
		}
	}

	return string(buf)
}
