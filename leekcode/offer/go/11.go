package _go

/*
剑指 Offer 11. 旋转数组的最小数字

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

示例 1：

输入：[3,4,5,1,2]
输出：1
示例 2：

输入：[2,2,2,0,1]
输出：0
*/

func minArray(numbers []int) int {
	if len(numbers) > 0 && numbers[0] < numbers[len(numbers)-1] {
		return numbers[0]
	}

	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			return numbers[i]
		}
	}

	//sort.Ints(numbers)
	//return numbers[0]
	return numbers[0]
}
