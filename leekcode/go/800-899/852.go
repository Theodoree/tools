package _00_899

/*
852. 山脉数组的峰顶索引

我们把符合下列属性的数组 A 称作山脉：

A.length >= 3
存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。
*/

func peakIndexInMountainArray(arr []int) int {

	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] > arr[left] {
			left++
		} else if arr[mid] > arr[right] {
			right--
		} else {
			return mid
		}
	}

	return mid

}
