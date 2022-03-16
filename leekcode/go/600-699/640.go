package _00_699


func solveEquation(equation string) string {
	var x, num, temp int
	minus, equal := 1, 1
	for i := range equation {
		if equation[i] >= '0' && equation[i] <= '9' {
			temp *= 10
			temp += int(equation[i] - '0')
		} else if equation[i] == 'x' {
			if temp == 0 {
				temp = 1
			}
			if i > 0 && equation[i-1] == '0' {
				temp = 0
			}
			x += temp * minus * equal
			temp = 0
			minus = 1
		} else if equation[i] == '=' {
			num += temp * minus * equal
			temp = 0
			minus = 1
			equal = -1
		} else if equation[i] == '-' {
			num += temp * minus * equal
			temp = 0
			minus = -1
		} else {
			num += temp * minus * equal
			temp = 0
			minus = 1
		}
	}
	num += temp * minus * equal
	//fmt.Println(x, num)
	if x == 0 && num == 0 {
		return "Infinite solutions"
	}
	if x == 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(-1*num/x)
}


