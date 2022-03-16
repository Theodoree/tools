package _100_2199


func addSpaces(s string, spaces []int) string {
	spaces = append(spaces, len(s)) // 小技巧：把 s 长度加到数组末尾，这样可以在循环内处理最后一段
	ans := []byte(s[:spaces[0]])
	for i := 1; i < len(spaces); i++ {
		ans = append(ans, ' ')
		ans = append(ans, s[spaces[i-1]:spaces[i]]...)
	}
	return string(ans)
}

