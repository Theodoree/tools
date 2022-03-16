package _go

/*
剑指 Offer II 033. 变位词组
给定一个字符串数组 strs ，将 变位词 组合在一起。 可以按任意顺序返回结果列表。
注意：若两个字符串中每个字符出现的次数都相同，则称它们互为变位词。
*/
func groupAnagrams(strs []string) [][]string {

	var m = make(map[string][]string, len(strs))
	for i := 0; i < len(strs); i++ {
		var buf [26]byte
		for _, v := range strs[i] {
			buf[v-'a']++
		}
		m[string(buf[:])] = append(m[string(buf[:])], strs[i])
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}
