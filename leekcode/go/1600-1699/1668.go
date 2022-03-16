package _600_1699

import (
	"fmt"
	"strings"
)

/*
1668. 最大重复子字符串
给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为 k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence 的子串，那么重复值 k 为 0 。
给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。
*/

func maxRepeating(sequence string, word string) int {
	for i := 0; i < len(sequence); i++ {
		fmt.Println(strings.Repeat(word, i+1))
		if strings.Index(sequence, strings.Repeat(word, i+1)) == -1 {
			return i
		}
	}
	return len(sequence) / len(word)
}
