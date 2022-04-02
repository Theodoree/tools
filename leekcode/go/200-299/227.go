package _00_299


/*
227. 基本计算器 II
给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
整数除法仅保留整数部分。
你可以假设给定的表达式总是有效的。所有中间结果将在 [-231, 231 - 1] 的范围内。
注意：不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
*/
func calculate(s string) int {
	var flag []byte
	var nums []int

	var i int
	for i < len(s) {
		switch {
		case s[i] == ' ':
		case s[i] >= '0':
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
