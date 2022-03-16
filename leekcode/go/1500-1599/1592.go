package _500_1599

/*
1592. 重新排列单词间的空格
给你一个字符串 text ，该字符串由若干被空格包围的单词组成。每个单词由一个或者多个小写英文字母组成，并且两个单词之间至少存在一个空格。题目测试用例保证 text 至少包含一个单词 。
请你重新排列空格，使每对相邻单词之间的空格数目都 相等 ，并尽可能 最大化 该数目。如果不能重新平均分配所有空格，请 将多余的空格放置在字符串末尾 ，这也意味着返回的字符串应当与原 text 字符串的长度相等。
返回 重新排列空格后的字符串 。
*/
func reorderSpaces(text string) string {
	var (
		word  []string
		space int
	)
	idx := -1
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			if idx != -1 {
				word = append(word, text[idx:i])
				idx = -1
			}
			space++
			continue
		} else if idx == -1 {
			idx = i
		}
	}
	if idx != -1 {
		word = append(word, text[idx:])
	}

	var space_str string
	count := space

	if len(word) > 1 {
		for i := 0; i < space/(len(word)-1); i++ {
			space_str += " "
		}
		count = space % (len(word) - 1)
	}
	text = strings.Join(word, space_str)

	for count > 0 {
		text += " "
		count--
	}
	return text

}
