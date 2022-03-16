package _00_299

import (
	"strings"
)

/*
211. 添加与搜索单词 - 数据结构设计

设计一个支持以下两种操作的数据结构：

void addWord(word)
bool search(word)
search(word) 可以搜索文字或正则表达式字符串，字符串只包含字母 . 或 a-z 。 . 可以表示任何一个字母。

示例:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
说明:

你可以假设所有单词都是由小写字母 a-z 组成的。
*/

type WordDictionary struct {
	root *Node
}
type Node struct {
	arr [26]*Node
	m   map[string]struct{}
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{root: &Node{
		arr: [26]*Node{},
		m:   nil,
	}}

}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {

	cur := this.root

	for k, v := range word {
		idx := v - 'a'
		if cur.arr[idx] != nil {
			cur = cur.arr[idx]
		} else {
			n := &Node{
				arr: [26]*Node{},
				m:   make(map[string]struct{}),
			}
			cur.arr[idx] = n
			cur = n
		}

		if k == len(word)-1 {
			cur.m[word] = struct{}{}
		}

	}

}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var ok bool

	if strings.Index(word, `.`) >= 0 {
		DFS(this.root, word, &ok)
	} else {

		cur := this.root

		for k, v := range word {
			idx := v - 'a'
			if cur.arr[idx] != nil {
				cur = cur.arr[idx]
			} else {
				return false
			}

			if k == len(word)-1 {
				_, ok = cur.m[word]
			}
		}
	}

	return ok
}

func DFS(n *Node, word string, ok *bool) {
	if len(word) == 0 || word == "" {
		if len(n.m) > 0 {
			*ok = true
		}
		return
	}

	for k, v := range n.arr {
		if v != nil {
			if word[0] == '.' {
				DFS(v, word[1:], ok)
			} else if word[0] == uint8(k+'a') {
				DFS(v, word[1:], ok)
			}
		}
	}
}
