package _100_2199

import "strings"

/*
2114. 句子中的最多单词数
一个 句子 由一些 单词 以及它们之间的单个空格组成，句子的开头和结尾不会有多余空格。
给你一个字符串数组 sentences ，其中 sentences[i] 表示单个 句子 。
请你返回单个句子里 单词的最多数目 。
*/
func mostWordsFound(sentences []string) int {
	var max int
	for i := 0; i < len(sentences); i++ {
		cnt := strings.Count(sentences[i], " ") + 1
		if cnt > max {
			max = cnt
		}
	}
	return max
}
