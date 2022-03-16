package _300_1399

// 穷举
func countLargestGroup(n int) int {
	// 统计
	var sum = 1

	sumMpCnt := [46]int{} // 最大99999 = 46
	for num := 1; num <= n; num++ {
		sumMpCnt[sum]++
		numCur := num
		for numCur%10 == 9 {
			sum -= 9
			numCur /= 10
		}
		sum++
	}

	// 最大数量
	var cnt, cntLargest int
	for sum := range sumMpCnt {
		if cnt < sumMpCnt[sum] {
			cnt = sumMpCnt[sum]
			cntLargest = 1
		} else if sumMpCnt[sum] == cnt {
			cntLargest++
		}
	}
	return cntLargest
}
