package _500_1599

/*
1556. 千位分隔数
给你一个整数 n，请你每隔三位添加点（即 "." 符号）作为千位分隔符，并将结果以字符串格式返回。
*/
func thousandSeparator(n int) string {

	if n == 0 {
		return "0"
	}

	var s string
	var count int

	for n > 0 {

		if count > 0 && count%3 == 0 {
			s = "." + s
		}
		cnt := n % 10
		n /= 10

		s = strconv.Itoa(cnt) + s
		count++

	}

	return s

}
