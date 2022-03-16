package _go

/*
面试题 16.26. 计算器
给定一个包含正整数、加(+)、减(-)、乘(*)、除(/)的算数表达式(括号除外)，计算其结果。
表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。
*/
func calculate(s string) int {
	var flag []byte
	var nums []int

	var i int
	for i < len(s) {
		switch {
		case s[i] == ' ':
		case s[i] <= '9':
			var n int
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				n = n*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, n)
			continue
		default:
			flag = append(flag, s[i])
		}
		i++
	}

	for len(flag) > 0 {

		switch flag[0] {
		case '*':
			nums[1] = nums[0] * nums[1]
		case '/':
			nums[1] = nums[0] / nums[1]
		case '+':
			if len(flag) > 1 {
				switch flag[1] {
				case '*':
					nums[2] = nums[1] * nums[2]
					flag[1] = flag[0]
					nums[1] = nums[0]
				case '/':
					nums[2] = nums[1] / nums[2]
					flag[1] = flag[0]
					nums[1] = nums[0]
				default:
					nums[1] = nums[0] + nums[1]
				}
			} else {
				nums[1] = nums[0] + nums[1]
			}

		case '-':
			if len(flag) > 1 {
				switch flag[1] {
				case '*':
					nums[2] = nums[1] * nums[2]
					flag[1] = flag[0]
					nums[1] = nums[0]
				case '/':
					nums[2] = nums[1] / nums[2]
					flag[1] = flag[0]
					nums[1] = nums[0]
				default:
					nums[1] = nums[0] - nums[1]
				}
			} else {
				nums[1] = nums[0] - nums[1]
			}
		}

		flag = flag[1:]
		nums = nums[1:]

	}

	return nums[0]
}
