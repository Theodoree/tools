package _00_499

/*
408. 有效单词缩写
字符串可以用 缩写 进行表示，缩写 的方法是将任意数量的 不相邻 的子字符串替换为相应子串的长度。例如，字符串 "substitution" 可以缩写为（不止这几种方法）：
*/
func validWordAbbreviation(word string, abbr string) bool {

	for len(abbr) > 0 {
		for len(abbr) > 0 && abbr[0]-'0' > 9 {
			if len(word) == 0 || abbr[0] != word[0] {
				return false
			}
			abbr = abbr[1:]
			word = word[1:]
		}

		if len(abbr) > 0 && abbr[0] == '0' { // 前置零
			return false
		}
		var cnt int
		for len(abbr) > 0 && abbr[0]-'0' <= 9 {
			cnt = cnt*10 + int(abbr[0]) - '0'
			abbr = abbr[1:]
		}

		if len(word) < cnt {
			return false
		}
		word = word[cnt:]

	}

	return len(word) == 0

}
