package _00_199


func atoi(i string) int {
	flag:=false
	if i[0] == '-'{
		i = i[1:]
		flag = true
	}
	var cnt int
	for len(i) > 0 {
		cnt = cnt*10 + int(i[0]-'0')
		i = i[1:]
	}
	if flag {
		return -cnt
	}
	return cnt
}

/*
150. 逆波兰表达式求值
根据 逆波兰表示法，求表达式的值。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
注意 两个整数之间的除法只保留整数部分。
可以保证给定的逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
*/
func evalRPN(tokens []string) int {
	var queue []byte
	var nums []int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+", "-", "*", "/":
			queue = append(queue, tokens[i][0])
			switch queue[len(queue)-1] {
			case '+':
				nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
			case '-':
				nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
			case '*':
				nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
			case '/':
				nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
			}
			nums = nums[:len(nums)-1]
			queue = queue[:len(queue)-1]
		default:
			nums = append(nums, atoi(tokens[i]))
		}
	}

	if len(queue) > 0 {
		switch queue[len(queue)-1] {
		case '+':
			nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
		case '-':
			nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
		case '*':
			nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
		case '/':
			nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
		}
		nums = nums[:len(nums)-1]
		queue = queue[:len(queue)-1]
	}

	return nums[0]
}

