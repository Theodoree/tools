package _go

/*
剑指 Offer 44. 数字序列中某一位的数字
数字以0123456789101112131415…的格式序列化到一个字符序列中。在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
*/
func findNthDigit(n int) int {
	var i int

	for n > i*int(math.Pow(10, float64(i-1))*9) {
		n -= i * int(math.Pow(10, float64(i-1))*9)
		i++
	}

	num := int(float64(n-1)/float64(i) + math.Pow(10, float64(i-1)))
	a := strconv.Itoa(num)
	if n%i == 0 {
		cnt, _ := strconv.Atoi(string(a[i-1]))
		return cnt
	}

	cnt, _ := strconv.Atoi(string(a[n%i-1]))
	return cnt
}
