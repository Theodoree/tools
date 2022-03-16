package _400_1499


/*
1422. 分割字符串的最大得分
给你一个由若干 0 和 1 组成的字符串 s ，请你计算并返回将该字符串分割成两个 非空 子字符串（即 左 子字符串和 右 子字符串）所能获得的最大得分。
「分割字符串的得分」为 左 子字符串中 0 的数量加上 右 子字符串中 1 的数量。
*/
func maxScore(s string) int {
	// 左边的0比1多,然后计算右边的1
	leftZero := 0
	rightOne := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0':
			leftZero++
		}
	}

	var max int

	for i := len(s) - 1; i > 0; i-- {
		switch s[i] {
		case '0':
			leftZero--
		case '1':
			rightOne++
		}

		if leftZero+rightOne > max {
			max = leftZero + rightOne
		}
	}
	return max

}