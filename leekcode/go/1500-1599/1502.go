package _500_1599

/*
1502. 判断能否形成等差数列
给你一个数字数组 arr 。
如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。
如果可以重新排列数组形成等差数列，请返回 true ；否则，返回 false 。
*/
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	sub := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] != sub {
			return false
		}
	}
	return true

}
