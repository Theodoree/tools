package _go

/*
面试题 10.05. 稀疏数组搜索
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。
*/
func findString(words []string, s string) int {
	for i := 0; i < len(words); i++ {
		if len(words[i]) == 0 {
			continue
		}
		if words[i][0] > s[0] {
			return -1
		}
		if words[i] == s {
			return i
		}
	}
	return -1
}
