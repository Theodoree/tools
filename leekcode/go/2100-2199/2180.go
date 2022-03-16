package _100_2199





/*
2180. 统计各位数字之和为偶数的整数个数
给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
*/
func countEven(num int) int {
	var cnt int
	odd:= func(i int) bool {
		var cnt int
		for i>0 {
			cnt+=i%10
			i/=10
		}
		return cnt%2 == 0
	}

	for i:=2;i<=num;i++{

		if odd(i){
			cnt++
		}
	}
	return cnt

}
