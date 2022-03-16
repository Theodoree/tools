package _500_1599

/*
1539. 第 k 个缺失的正整数
给你一个 严格升序排列 的正整数数组 arr 和一个整数 k 。

请你找到这个数组里第 k 个缺失的正整数。

示例 1：

输入：arr = [2,3,4,7,11], k = 5
输出：9
解释：缺失的正整数包括 [1,5,6,8,9,10,12,13,...] 。第 5 个缺失的正整数为 9 。
示例 2：

输入：arr = [1,2,3,4], k = 2
输出：6
解释：缺失的正整数包括 [5,6,7,...] 。第 2 个缺失的正整数为 6 。

*/

func findKthPositive(arr []int, k int) int {
	var idx int
	var miss int
	for i := 1; i <= arr[len(arr)-1]+k; i++ {
		if idx < len(arr) && i == arr[idx] {
			idx++
			continue
		}

		miss++
		if miss == k {
			return i
		}
	}

	return 0
}
