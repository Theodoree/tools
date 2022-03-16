package _go

import "sort"

/*
面试题 17.15. 最长单词
给定一组单词words，编写一个程序，找出其中的最长单词，且该单词由这组单词中的其他单词组合而成。若有多个长度相同的结果，返回其中字典序最小的一项，若没有符合要求的单词则返回空字符串。

示例：

输入： ["cat","banana","dog","nana","walk","walker","dogwalker"]
输出： "dogwalker"
解释： "dogwalker"可由"dog"和"walker"组成。
*/
type trie struct {
	child [26]*trie
	end   bool
}

func (t *trie) Insert(s string) {
	cur := t
	for _, v := range s {
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &trie{}
		}
		cur = cur.child[v-'a']
	}
	cur.end = true
}

func (t *trie) search(s string,flag bool) bool{

	cur:=t
	for idx,v:=range s{
		if cur.child[v-'a'] == nil {
			return false
		}
		cur = cur.child[v-'a']
		if cur.end {
			if flag && idx == len(s)-1{
				return true
			}
			if t.search(s[idx+1:],true) {
				return true
			}
		}
	}


	return false
}


func longestWord(words []string) string {
	var t trie
	for _,v:=range words{
		t.Insert(v)
	}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) > len(words[j])
	})


	for i:=0;i<len(words);i++{
		if t.search(words[i],false) {
			return words[i]
		}
	}


	return ""
}

