package _go

/*
面试题 08.04. 幂集
幂集。编写一种方法，返回某集合的所有子集。集合中不包含重复的元素。
说明：解集不能包含重复的子集。
*/
func subsets(nums []int) [][]int {
	var result [][]int
	var l uint = uint(len(nums))
	var n = 1 << l
	for i := 0; i < n; i++ {
		var cur []int
		for j := 0; j < int(l); j++ {
			if i&(1<<uint(j)) == 0 {
				cur = append(cur, nums[j])
			}
		}
		result = append(result, cur)
	}

	return result
}
