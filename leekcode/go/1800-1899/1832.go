package _800_1899

/*
1832. 判断句子是否为全字母句
全字母句 指包含英语字母表中每个字母至少一次的句子。

给你一个仅由小写英文字母组成的字符串 sentence ，请你判断 sentence 是否为 全字母句 。

如果是，返回 true ；否则，返回 false 。
*/
func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	alphabet := [26]byte{}

	for _, v := range sentence {
		alphabet[v-'a']++
	}

	for _, v := range alphabet {
		if v == 0 {
			return false
		}
	}

	return true
}
