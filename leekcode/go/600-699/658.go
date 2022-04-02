package _00_699


/*
658. 找到 K 个最接近的元素
给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的
*/
func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) <= k {
		return arr
	}
	left, right := 0, len(arr)-k
	for left < right {
		mid := left + (right-left)>>2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
