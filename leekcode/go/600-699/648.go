package _00_699

import "strings"

/*
剑指 Offer II 063. 替换单词
在英语中，有一个叫做 词根(root) 的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为 继承词(successor)。例如，词根an，跟随着单词 other(其他)，可以形成新的单词 another(另一个)。
现在，给定一个由许多词根组成的词典和一个句子，需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
需要输出替换之后的句子。
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

func (t *trie) exits(s string) string {
	cur := t

	for i := 0; i < len(s); i++ {
		if cur.end {
			return s[:i]
		}
		cur = cur.child[s[i]-'a']
		if cur == nil {
			return s
		}

	}
	return s
}

func replaceWords(dictionary []string, sentence string) string {
	var t trie
	for _, v := range dictionary {
		t.Insert(v)
	}

	var result []byte
	result = make([]byte, 0, len(sentence))
	for start, end := 0, 0; end <= len(sentence); end++ {
		if end == len(sentence) || sentence[end] == ' ' {
			result = append(result, t.exits(sentence[start:end])...)
			result = append(result, ' ')
			end++
			start = end
		}
	}
	return string(result[:len(result)-1])
}

func replaceWords(dictionary []string, sentence string) string {
	if len(dictionary) == 0 {
		return sentence
	}
	str := strings.Split(sentence, " ")
	dicmap := make(map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		dicmap[dictionary[i]] = true
	}
	for i := 0; i < len(str); i++ {
		for j := 1; j < len(str[i]); j++ {
			if _, ok := dicmap[str[i][:j]]; ok {
				str[i] = str[i][:j]
			}
		}
	}
	res := strings.Join(str, " ")
	return res
}
