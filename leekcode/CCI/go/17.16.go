package _go

/*
面试题 17.16. 按摩师
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。
注意：本题相对原题稍作改动



*/
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := func(i, j int) int {
		if j > i {
			return j
		}
		return i
	}
	dp0 := 0
	dp1 := nums[0]

	for i := 1; i < len(nums); i++ {
		//fmt.Println(dp0, dp1, i-1)
		tdp0 := max(dp0, dp1)
		tdp1 := dp0 + nums[i]

		dp0 = tdp0
		dp1 = tdp1
	}

	return max(dp0, dp1)

}
