package _00_299

/*
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。如果可以变为 1，那么这个数就是快乐数。
质数
*/
func isHappy(n int) bool {

	if n == 1 {
		return true
	}
	if n == 4 {
		return false
	}

	s := GetSlice(n)

	var sum int
	for i := 0; i < len(s); i++ {
		sum += s[i] * s[i]
	}
	return isHappy(sum)
}

func GetSlice(n int) []int {
	var result []int
	var cnt int
	for n != 0 {
		cnt = n % 10
		result = append(result, cnt)
		n = n / 10
	}

	return result
}
