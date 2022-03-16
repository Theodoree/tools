package _900_1999

/*
1935. 可以输入的最大单词数
键盘出现了一些故障，有些字母键无法正常工作。而键盘上所有其他键都能够正常工作。
给你一个由若干单词组成的字符串 text ，单词间由单个空格组成（不含前导和尾随空格）；另有一个字符串 brokenLetters ，由所有已损坏的不同字母键组成，返回你可以使用此键盘完全输入的 text 中单词的数目。
*/
func canBeTypedWords(text string, brokenLetters string) int {
	var buf [26]int
	for _, v := range brokenLetters {
		buf[v-'a'] = -1
	}
	var count int
	var flag bool
	for _, v := range text {
		if v == ' ' {
			if !flag {
				count++
			}
			flag = false
			continue
		}
		if flag {
			continue
		}

		if buf[v-'a'] == -1 {
			flag = true
		}
	}

	if !flag {
		count++
	}

	return count
}
