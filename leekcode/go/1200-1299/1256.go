package _200_1299



/*
1256. 加密数字
给你一个非负整数 num ，返回它的「加密字符串」。
加密的过程是把一个整数用某个未知函数进行转化，你需要从下表推测出该转化函数：
*/
func encode(num int) string {
	num++
	var result string
	for num > 1 {
		if num&1 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}
		num >>= 1
	}

	return result

}

