package __99

/*
49. 字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
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
