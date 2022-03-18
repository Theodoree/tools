package _00_399

/*
388. 文件的最长绝对路径
假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：
*/
func lengthLongestPath(input string) int {
	rows := strings.Split(input, "\n")
	res := 0
	buff := []int{}
	// fmt.Println(rows)
	for _, row := range rows {
		depth := 0
		for i := 0; i < len(row); i++ {
			if row[i] != '\t' {
				break
			}
			depth++
		}
		if len(buff) <= depth {
			buff = append(buff, len(row)-depth)
		} else {
			buff[depth] = len(row) - depth
			buff = buff[:depth+1]
		}
		if !strings.Contains(row, ".") {
			continue
		}
		sum := 0
		for i := 0; i < len(buff); i++ {
			sum += buff[i]
		}
		sum += len(buff) - 1
		res = max(res, sum)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
