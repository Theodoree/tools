package _go


/*
面试题 17.10. 主要元素
数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。

示例 1：

输入：[1,2,5,9,5,9,5,5,5]
输出：5


示例 2：

输入：[3,2]
输出：-1


示例 3：

输入：[2,2,1,1,1,2,2]
输出：2

*/
func majorityElement(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}

	return -1
}

