package _900_5999

/*
5942. 找出 3 位偶数
给你一个整数数组 digits ，其中每个元素是一个数字（0 - 9）。数组中可能存在重复元素。
你需要找出 所有 满足下述条件且 互不相同 的整数：
该整数由 digits 中的三个元素按 任意 顺序 依次连接 组成。
该整数不含 前导零
该整数是一个 偶数
例如，给定的 digits 是 [1, 2, 3] ，整数 132 和 312 满足上面列出的全部条件。
将找出的所有互不相同的整数按 递增顺序 排列，并以数组形式返回。
*/
func findEvenNumbers(digits []int) []int {
	var max int
	var buf [10]int
	for _, v := range digits {
		buf[v]++
		if v > max {
			max = v
		}
	}
	max = max*100 + max*10 + max

	var result []int
	var t [10]int
	for i := 100; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		cur := i
		for cur > 0 {
			t[cur%10]++
			cur /= 10
		}
		flag := true
		for num, v := range t {
			if v > 0 && buf[num] < v {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, i)
		}
		t = [10]int{}

	}

	return result

}
