package _900_1999

/*
1957. 删除字符使字符串变好
一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。
给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。
请你返回删除后的字符串。题目数据保证答案总是 唯一的 。
*/
func makeFancyString(s string) string {
	val := make([]byte, 0, 10)
	var lastByte byte
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == lastByte && count >= 3 {
			count++
			continue
		}
		if s[i] != lastByte {
			lastByte = s[i]
			count = 1
		}
		lastByte = s[i]
		val = append(val, lastByte)
		count++
	}
	return string(val)
}
