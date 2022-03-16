package _200_1299

/*
1271. 十六进制魔术数字
你有一个十进制数字，请按照此规则将它变成「十六进制魔术数字」：首先将它变成字母大写的十六进制字符串，然后将所有的数字 0 变成字母 O ，将数字 1  变成字母 I 。
如果一个数字在转换后只包含 {"A", "B", "C", "D", "E", "F", "I", "O"} ，那么我们就认为这个转换是有效的。
给你一个字符串 num ，它表示一个十进制数 N，如果它的十六进制魔术数字转换是有效的，请返回转换后的结果，否则返回 "ERROR" 。
*/
func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	var result []byte
	for n > 0 {
		cnt := n & 15
		n >>= 4
		if cnt >= 2 && cnt <= 9 {
			return "ERROR"
		}

		switch cnt {
		case 0:
			result = append(result, 'O')
		case 1:
			result = append(result, 'I')
		default:
			result = append(result, 'A'+byte(cnt%10))
			//case 10:
			//	result = append(result,'A')
			//case 11:
			//	result = append(result,'B')
			//case 12:
			//	result = append(result,'C')
			//case 13:
			//	result = append(result,'D')
			//case 14:
			//	result = append(result,'E')
			//case 15:
			//	result = append(result,'F')
		}

	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}

	return string(result)

}
