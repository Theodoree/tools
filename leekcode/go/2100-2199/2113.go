package _100_2199

/*
2113. 查询删除和添加元素后的数组
给你一个 下标从 0 开始 的数组 nums。一开始，在第 0 分钟，数组没有变化。此后每过一分钟，数组的 最左边 的元素将被移除，直到数组为空。然后，每过一分钟，数组的 尾部 将添加一个元素，添加的顺序和删除的顺序相同，直到数组被复原。此后上述操作无限循环进行。
举个例子，如果 nums = [0, 1, 2]，那么数组将按如下流程变化：[0,1,2] → [1,2] → [2] → [] → [0] → [0,1] → [0,1,2] → [1,2] → [2] → [] → [0] → [0,1] → [0,1,2] → ...
然后给你一个长度为 n 的二维数组 queries，其中 queries[j] = [timej, indexj]，表示第 j 个查询。第 j 个查询的答案定义如下：
如果在时刻 timej，indexj < nums.length，那么答案是此时的 nums[indexj]；
如果在时刻 timej，indexj >= nums.length，那么答案是 -1。
请返回一个长度为 n 的整数数组 ans，其中 ans[j] 为第 j 个查询的答案。
*/

func elementInNums(nums []int, queries [][]int) []int {

	var result []int
	for i := 0; i < len(queries); i++ {
		t := queries[i][0]
		idx := queries[i][1]

		if t/len(nums) %2 ==  0 {
			n := t % len(nums)
			if n+idx >= len(nums) {
				result = append(result, -1)
			} else {
				result = append(result, nums[n+idx])
			}
		} else {
			n := t % len(nums)
			if n <= idx {
				result = append(result, -1)
			} else {
				result = append(result, nums[idx])
			}
		}

	}

	return result
}
