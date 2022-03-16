package _100_1199

import "sort"

/*
1196. 最多可以买到的苹果数量
楼下水果店正在促销，你打算买些苹果，arr[i] 表示第 i 个苹果的单位重量。
你有一个购物袋，最多可以装 5000 单位重量的东西，算一算，最多可以往购物袋里装入多少苹果。
*/

func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)
	var sum int
	var count int
	for _, v := range weight {
		sum += v
		if sum > 5000 {
			break
		}
		count++
	}

	return count

}
