package _00_799

import "sort"

/*
720. 词典中最长的单词
给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。

若无答案，则返回空字符串。



示例 1：

输入：
words = ["w","wo","wor","worl", "world"]
输出："world"
解释：
单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。
示例 2：

输入：
words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
输出："apple"
解释：
"apply"和"apple"都能由词典中的单词组成。但是"apple"的字典序小于"apply"。


提示：

所有输入的字符串都只包含小写字母。
words数组长度范围为[1,1000]。
words[i]的长度范围为[1,30]。
通过次数21,459提交次数44,323
*/
type trie struct {
	root node
}

type node struct {
	node [26]*node
	end  bool
}

func (t *trie) Insert(s string) {
	cur := &t.root
	for _, v := range s {
		if cur.node[v-'a'] == nil {
			cur.node[v-'a'] = &node{}
		}
		cur = cur.node[v-'a']
	}
	cur.end = true
}

func (t *trie) Check(s string) bool {
	cur := &t.root
	for i := 0; i < len(s)-1; i++ {
		cur = cur.node[s[i]-'a']
		if cur == nil || !cur.end {
			return false
		}
	}
	return true
}

func longestWord(words []string) string {
	var t trie
	for _, v := range words {
		t.Insert(v)
	}

	sort.Slice(words, func(i, j int) bool {
		l1, l2 := len(words[i]), len(words[j])
		if l1 == l2 {
			return words[i] < words[j]
		}
		return l1 > l2
	})

	for i := 0; i < len(words); i++ {
		if t.Check(words[i]) {
			return words[i]
		}
	}

	return ""
}
