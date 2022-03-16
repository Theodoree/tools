package _600_1699

/*
1694. 重新格式化电话号码
给你一个字符串形式的电话号码 number 。number 由数字、空格 ' '、和破折号 '-' 组成。
请你按下述方式重新格式化电话号码。
首先，删除 所有的空格和破折号。
其次，将数组从左到右 每 3 个一组 分块，直到 剩下 4 个或更少数字。剩下的数字将按下述规定再分块：
2 个数字：单个含 2 个数字的块。
3 个数字：单个含 3 个数字的块。
4 个数字：两个分别含 2 个数字的块。
最后用破折号将这些块连接起来。注意，重新格式化过程中 不应该 生成仅含 1 个数字的块，并且 最多 生成两个含 2 个数字的块。
返回格式化后的电话号码。
*/
func reformatNumber(number string) string {

	str := make([]byte, 0, len(number))
	for _, v := range number {
		if v == '-' || v == ' ' {
			continue
		}
		str = append(str, byte(v))
	}

	if len(str) < 3 {
		return string(str)
	}

	var result = make([]byte, 0, len(number))

	for i := 0; i < len(str)-len(str)%3; i++ {
		if i > 0 && i%3 == 0 {
			result = append(result, '-')
		}
		result = append(result, str[i])
	}
	switch len(str) % 3 {
	case 2:
		result = append(result, '-')
		result = append(result, str[len(str)-2:]...)
	case 1:
		result = result[:len(result)-3]
		result = append(result, str[len(str)-4:len(str)-2]...)
		result = append(result, '-')
		result = append(result, str[len(str)-2:]...)

	}

	return string(result)
}
