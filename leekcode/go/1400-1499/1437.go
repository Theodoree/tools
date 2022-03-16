package _400_1499

/*
1437. 是否所有 1 都至少相隔 k 个元素
给你一个由若干 0 和 1 组成的数组 nums 以及整数 k。如果所有 1 都至少相隔 k 个元素，则返回 True ；否则，返回 False 。
*/

func kLengthApart(nums []int, k int) bool {
	var n, n1 int

	n = -1
	n1 = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != 1 {
			continue
		}

		switch {
		case n == -1:
			n = i
			continue
		case n1 == -1:
			n1 = i
			continue
		}

		if n1-n-1 < k {
			return false
		}
		n = n1
		n1 = i
	}
	if n1 != -1 && n != -1 && n1-n-1 < k {
		return false
	}
	return true

}
