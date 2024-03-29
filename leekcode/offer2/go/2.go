package _go



/*
剑指 Offer II 002. 二进制加法
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。



示例 1:

输入: a = "11", b = "10"
输出: "101"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"


提示：

每个字符串仅由字符 '0' 或 '1' 组成。
1 <= a.length, b.length <= 10^4
字符串如果不是 "0" ，就都不含前导零。
*/
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		b, a = a, b
	}

	buf := make([]byte, len(b)+1)

	for i := 0; i < len(b); i++ {
		var n byte
		if len(a)-i-1 >= 0 {
			n = a[len(a)-i-1] - '0'
		}

		buf[len(buf)-i-1] = b[len(b)-i-1] + n
	}

	var plus bool

	for i := 0; i < len(buf); i++ {
		switch buf[len(buf)-i-1] {
		case '0':
			if !plus {
				continue
			}

			buf[len(buf)-i-1] = '1'
			plus = false
		case '1':
			if !plus {
				continue
			}
			buf[len(buf)-i-1] = '0'
		case '2':
			if !plus {
				plus = true
				buf[len(buf)-i-1] = '0'
				continue
			}
			plus = true
			buf[len(buf)-i-1] = '1'
		}

	}

	if plus {
		buf[0] = '1'
	} else {
		buf = buf[1:]
	}

	return string(buf)

}

