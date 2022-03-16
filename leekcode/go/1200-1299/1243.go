package _200_1299

/*
1243. 数组变换
首先，给你一个初始数组 arr。然后，每天你都要根据前一天的数组生成一个新的数组。
第 i 天所生成的数组，是由你对第 i-1 天的数组进行如下操作所得的：
假如一个元素小于它的左右邻居，那么该元素自增 1。
假如一个元素大于它的左右邻居，那么该元素自减 1。
首、尾元素 永不 改变。
过些时日，你会发现数组将会不再发生变化，请返回最终所得到的数组。
*/
func transformArray(arr []int) []int {

	sum := make([]int, len(arr))
	for i := 1; i < len(arr)-1; i++ {

		for j := 1; j < len(arr)-1; j++ {
			if arr[j] < arr[j-1] && arr[j] < arr[j+1] {
				sum[j]++
			} else if arr[j] > arr[j-1] && arr[j] > arr[j+1] {
				sum[j]--
			}
		}
		for idx := range sum {
			arr[idx] += sum[idx]
			sum[idx] = 0
		}

	}
	return arr
}
