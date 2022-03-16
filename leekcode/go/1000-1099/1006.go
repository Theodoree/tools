package _000_1099



func clumsy(n int) (ans int) {
	stk := []int{n}
	n--

	index := 0 // 用于控制乘、除、加、减
	for n > 0 {
		switch index % 4 {
		case 0:
			stk[len(stk)-1] *= n
		case 1:
			stk[len(stk)-1] /= n
		case 2:
			stk = append(stk, n)
		default:
			stk = append(stk, -n)
		}
		n--
		index++
	}

	// 累加栈中数字
	for _, v := range stk {
		ans += v
	}
	return
}

