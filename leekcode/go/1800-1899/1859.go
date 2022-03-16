package _800_1899

/*
1859. 将句子排序
一个 句子 指的是一个序列的单词用单个空格连接起来，且开头和结尾没有任何空格。每个单词都只包含小写或大写英文字母。
我们可以给一个句子添加 从 1 开始的单词位置索引 ，并且将句子中所有单词 打乱顺序 。
比方说，句子 "This is a sentence" 可以被打乱顺序得到 "sentence4 a3 is2 This1" 或者 "is2 sentence4 This1 a3" 。
给你一个 打乱顺序 的句子 s ，它包含的单词不超过 9 个，请你重新构造并得到原本顺序的句子。
*/
func sortSentence(s string) string {

	var result [10]string

	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result[s[i]-'0'-1] = s[idx:i]
			idx = i + 2
		}
	}

	s = ""
	for _, v := range result {
		if len(v) == 0 {
			break
		}
		s += v + " "
	}
	s = s[:len(s)-1]

	return s
}
