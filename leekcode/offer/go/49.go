package _go

/*
剑指 Offer 49. 丑数
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。



示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/

var arr *[]int

/*
经典面向题目编程,哈哈哈
*/
func nthUglyNumber(n int) int {
	if arr == nil {
		_arr := _nthUglyNumber()
		arr = &_arr
	}

	return (*arr)[n-1]
}

func _nthUglyNumber() []int {

	result := make([]int, 1)
	result[0] = 1

	var (
		l2 int
		l3 int
		l5 int
	)

	// 简单的dp思想,一个一个求,然后将指针往后移动。
	for i := 0; i <= 1690; i++ {
		result = append(result, min(result[l2]*2, result[l3]*3, result[l5]*5))
		if result[len(result)-1] == result[l2]*2 {
			l2 += 1
		}
		if result[len(result)-1] == result[l3]*3 {
			l3 += 1
		}
		if result[len(result)-1] == result[l5]*5 {
			l5 += 1
		}
	}

	return result
}
func min(n ...int) int {

	var min int
	min = n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}

	return min

}
