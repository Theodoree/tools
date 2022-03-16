package _400_1499

/*
1408. 数组中的字符串匹配
给你一个字符串数组 words ，数组中的每个字符串都可以看作是一个单词。请你按 任意 顺序返回 words 中是其他单词的子字符串的所有单词。
如果你可以删除 words[j] 最左侧和/或最右侧的若干字符得到 word[i] ，那么字符串 words[i] 就是 words[j] 的一个子字符串。
*/
func stringMatching(words []string) []string {
	var result []string
	var buf [26][]int
	for i := 0; i < len(words); i++ {
		str := words[i]
		buf[str[0]-'a'] = append(buf[str[0]-'a'], i)
	}

	for i := 0; i < len(words); i++ {
		str := words[i]

		for idx, val := range str {
			for bufIdx, wordIdx := range buf[val-'a'] {
				if wordIdx == -1 || idx+len(words[wordIdx]) > len(str) || words[wordIdx] == str {
					continue
				}
				subStr := words[wordIdx]
				if subStr == str[idx:idx+len(subStr)] {
					result = append(result, subStr)
					buf[val-'a'][bufIdx] = -1

				}

			}
		}
	}
	return result

}
