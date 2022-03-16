package _900_1999

import "strings"

/*
1967. 作为子字符串出现在单词中的字符串数目
给你一个字符串数组 patterns 和一个字符串 word ，统计 patterns 中有多少个字符串是 word 的子字符串。返回字符串数目。
子字符串 是字符串中的一个连续字符序列。
*/
func numOfStrings(patterns []string, word string) int {
	var cnt int
	for _, v := range patterns {
		if strings.Index(word, v) >= 0 {
			cnt++
		}
	}
	return cnt
}
