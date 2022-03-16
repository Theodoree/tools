package _100_2199

/*
2119. 反转两次的数字
反转 一个整数意味着倒置它的所有位。

例如，反转 2021 得到 1202 。反转 12300 得到 321 ，不保留前导零 。
给你一个整数 num ，反转 num 得到 reversed1 ，接着反转 reversed1 得到 reversed2 。如果 reversed2 等于 num ，返回 true ；否则，返回 false 。
*/
func isSameAfterReversals(num int) bool { // 直接模拟就完事了
	cur := num
	var cnt int
	for cur > 0 {
		t := cur % 10
		cur /= 10
		cnt = cnt*10 + t
	}

	for cnt > 0 {
		t := cnt % 10
		cnt /= 10
		cur = cur*10 + t
	}

	return cur == num
}
