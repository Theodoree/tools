package _go

/*
面试题 17.11. 单词距离
有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
*/
func findClosest(words []string, word1 string, word2 string) int {
	var arr [2][]int
	for idx, v := range words {
		if v == word1 {
			arr[0] = append(arr[0], idx)
		}
		if v == word2 {
			arr[1] = append(arr[1], idx)
		}
	}

	var min int
	min = math.MaxInt64
	for i := 0; i < len(arr[0]) && min != 1; i++ {
		for j := 0; j < len(arr[1]); j++ {
			val := abs(arr[0][i] - arr[1][j])
			if val < min {
				min = val
			}

		}
	}

	return min
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

