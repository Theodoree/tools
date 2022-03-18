package _go


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
剑指 Offer II 036. 后缀表达式
根据 逆波兰表示法，求该后缀表达式的计算结果。
有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
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
