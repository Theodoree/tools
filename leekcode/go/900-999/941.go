package _00_999

/*
941. 有效的山脉数组
给定一个整数数组 arr，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

arr.length >= 3
在 0 < i < arr.length - 1 条件下，存在 i 使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]






示例 1：

输入：arr = [2,1]
输出：false
示例 2：

输入：arr = [3,5,5]
输出：false
示例 3：

输入：arr = [0,3,2,1]
输出：true


提示：

1 <= arr.length <= 104
0 <= arr[i] <= 104
*/
func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}
	left := 0
	right := len(arr) - 1
	var mid = 0
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] > arr[left] {
			left++
			if arr[left-1] >= arr[left] {
				return false
			}
		} else if arr[mid] > arr[right] {
			if arr[right-1] <= arr[right] {
				return false
			}
			right--
		} else {
			break
		}
	}
	if mid == 0 {
		return false
	}

	return arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1]
}
