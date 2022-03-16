package _200_1299

/*
1228. 等差数列中缺失的数字
有一个数组，其中的值符合等差数列的数值规律，也就是说：
在 0 <= i < arr.length - 1 的前提下，arr[i+1] - arr[i] 的值都相等。
我们会从该数组中删除一个 既不是第一个 也 不是最后一个的值，得到一个新的数组  arr。
给你这个缺值的数组 arr，请你帮忙找出被删除的那个数。
*/
func missingNumber(arr []int) int {
	start := arr[0]
	cnt := (arr[len(arr)-1] - start) / len(arr)
	if cnt == 0 {
		return arr[0]
	}
	arr = arr[1:]
	for len(arr) > 0 {
		if start+cnt != arr[0] {
			return start + cnt
		}
		start += cnt
		arr = arr[1:]

	}
	return 0
}
