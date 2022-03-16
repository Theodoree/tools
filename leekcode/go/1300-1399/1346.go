package _300_1399

/*
1346. 检查整数及其两倍数是否存在
给你一个整数数组 arr，请你检查是否存在两个整数 N 和 M，满足 N 是 M 的两倍（即，N = 2 * M）。

更正式地，检查是否存在两个下标 i 和 j 满足：

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]
*/
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {
			if arr[i]*2 == arr[j] || (arr[i] == arr[j]*2) {
				return true
			}
			if arr[i] >= 0 && arr[i]*2 < arr[j] {
				break
			}
		}
	}

	return false
}
