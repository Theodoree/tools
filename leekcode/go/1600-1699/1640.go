package _600_1699

/*
1640. 能否连接形成数组
给你一个整数数组 arr ，数组中的每个整数 互不相同 。另有一个由整数数组构成的数组 pieces，其中的整数也 互不相同 。请你以 任意顺序 连接 pieces 中的数组以形成 arr 。但是，不允许 对每个数组 pieces[i] 中的整数重新排序。
如果可以连接 pieces 中的数组形成 arr ，返回 true ；否则，返回 false 。
*/

func canFormArray(arr []int, pieces [][]int) bool {

	var count int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(pieces); j++ {
			if pieces[j] == nil || arr[i] != pieces[j][0] {
				continue
			}
			for i < len(arr) && arr[i] == pieces[j][0] {
				pieces[j] = pieces[j][1:]
				if len(pieces[j]) == 0 {
					break
				}
				i++
			}
			if len(pieces[j]) != 0 {
				return false
			}
			count++
			pieces[j] = nil
		}
	}

	return len(pieces) == count
}
