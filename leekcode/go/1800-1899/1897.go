package _800_1899

/*
1897. 重新分配字符使所有字符串都相等
给你一个字符串数组 words（下标 从 0 开始 计数）。
在一步操作中，需先选出两个 不同 下标 i 和 j，其中 words[i] 是一个非空字符串，接着将 words[i] 中的 任一 字符移动到 words[j] 中的 任一 位置上。
如果执行任意步操作可以使 words 中的每个字符串都相等，返回 true ；否则，返回 false 。

"可以执行无数次!!!!"
*/
func makeEqual(words []string) bool {

	var buf [26]int
	for _, v := range words {
		for _, val := range v {
			buf[val-'a']++
		}
	}

	for i := 0; i < len(buf); i++ {
		if buf[i] > 0 && buf[i]%len(words) != 0 {
			return false
		}
	}
	return true

}
