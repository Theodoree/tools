package _700_1799

/*
1768. 交替合并字符串
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
返回 合并后的字符串 。
*/
func mergeAlternately(word1 string, word2 string) string {
	sum := make([]byte, 0, len(word1)+len(word2))
	var flag bool
	for len(word1) > 0 && len(word2) > 0 {

		if !flag {
			sum = append(sum, word1[0])
			word1 = word1[1:]
			flag = true
			if len(word1) == 0 {
				sum = append(sum, word2...)
				return string(sum)
			}
		} else {
			sum = append(sum, word2[0])
			word2 = word2[1:]
			flag = false
			if len(word2) == 0 {
				sum = append(sum, word1...)
				return string(sum)
			}
		}
	}

	return string(sum)
}
