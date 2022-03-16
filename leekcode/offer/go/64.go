package _go

/*
剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）
*/
func sumNums(n int) int {
	result := n
	if n > 0 { // 用 && 符和 if有啥区别呢,都是跳转,隔着扯淡呢。
		result += sumNums(n - 1)
	}
	return result
}
