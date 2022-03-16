package _400_1499

/*
1426. 数元素
给你一个整数数组 arr， 对于元素 x ，只有当 x + 1 也在数组 arr 里时，才能记为 1 个数。
如果数组 arr 里有重复的数，每个重复的数单独计算。
*/
func countElements(arr []int) int {
	var m [1002]int
	var max int
	var min int
	min = math.MaxInt64
	for _, v := range arr {
		m[v]++
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	var result int
	for i := min; i <= max; i++ {
		if m[i] > 0 && m[i+1] > 0 {
			result += m[i]
		}
	}

	return result
}
