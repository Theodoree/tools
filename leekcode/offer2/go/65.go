package _go

import "sort"

/*
剑指 Offer II 065. 最短的单词编码
单词数组 words 的 有效编码 由任意助记字符串 s 和下标数组 indices 组成，且满足：
words.length == indices.length
助记字符串 s 以 '#' 字符结尾
对于每个下标 indices[i] ，s 的一个从 indices[i] 开始、到下一个 '#' 字符结束（但不包括 '#'）的 子字符串 恰好与 words[i] 相等
给定一个单词数组 words ，返回成功对 words 进行编码的最小助记字符串 s 的长度 。
*/
type trie struct {
	child [26]*trie
	end   bool
}

func (t *trie) suffix(s string) bool {
	cur := t
	var ok bool
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &trie{}
			ok = true
		}
		cur = cur.child[v-'a']
	}
	cur.end = true
	return ok
}
func minimumLengthEncoding(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	var t trie
	var sum int
	for _, v := range words {
		if t.suffix(v) {
			sum += len(v) + 1
		}
	}

	return sum
}

func minimumLengthEncoding(words []string) int {
	m := make(map[string]struct{}, 10)
	res := 0
	for _, w := range words {
		m[w] = struct{}{}
	}
	for w := range m {
		for i := 1; i < len(w); i++ {
			delete(m, w[i:])
		}
	}
	for w := range m {
		res += 1 + len(w)
	}
	return res
}
