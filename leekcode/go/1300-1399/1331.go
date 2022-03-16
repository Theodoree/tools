package _300_1399

import "sort"

/*
1331. 数组序号转换
给你一个整数数组 arr ，请你将数组中的每个元素替换为它们排序后的序号。

序号代表了一个元素有多大。序号编号的规则如下：

序号从 1 开始编号。
一个元素越大，那么序号越大。如果两个元素相等，那么它们的序号相同。
每个数字的序号都应该尽可能地小。
*/
func arrayRankTransform(arr []int) []int {
	hashMap := make(map[int]int, len(arr))
	for i := range arr {
		hashMap[arr[i]] = 0
	}
	copyArr := make([]int, len(hashMap))
	count := 0
	for val := range hashMap {
		copyArr[count] = val
		count++
	}
	sort.Ints(copyArr)
	for i := range copyArr {
		hashMap[copyArr[i]] = i + 1
	}
	for i := range arr {
		arr[i] = hashMap[arr[i]]
	}
	return arr
}
