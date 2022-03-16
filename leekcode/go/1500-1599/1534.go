package _500_1599

/*
1534. 统计好三元组
给你一个整数数组 arr ，以及 a、b 、c 三个整数。请你统计其中好三元组的数量。

如果三元组 (arr[i], arr[j], arr[k]) 满足下列全部条件，则认为它是一个 好三元组 。
0 <= i < j < k < arr.length
|arr[i] - arr[j]| <= a
|arr[j] - arr[k]| <= b
|arr[i] - arr[k]| <= c
其中 |x| 表示 x 的绝对值。
*/
func countGoodTriplets(arr []int, a int, b int, c int) int {
	var count int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) > b || (abs(arr[i]-arr[k]) > c) {
					continue
				}
				count++
			}
		}
	}
	return count
}
func abs(i int) int {
	if i < 0 {
		i *= -1
	}
	return i
}
