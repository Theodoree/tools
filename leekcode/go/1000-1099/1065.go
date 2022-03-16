package _000_1099

type trie struct {
	child [26]*trie
	end   bool
}

func (t *trie) Insert(s string) bool {
	cur := t
	var ok bool
	for _, v := range s {
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &trie{}
			ok = true
		}
		cur = cur.child[v-'a']
	}
	cur.end = true
	return ok
}
func (t *trie) exits(s string, m map[string]bool) (int, bool) {
	cur := t

	var i int
	for i = 0; i < len(s); i++ {
		if cur.end && !m[s[:i]] {
			m[s[:i]] = true
			return i, true
		}
		cur = cur.child[s[i]-'a']
		if cur == nil {
			return 0, false

		}
	}
	if cur.end && !m[s[:i]] {
		m[s[:i]] = true
		return i, true
	}
	return 0, false
}

/*
1065. 字符串的索引对
给出 字符串 text 和 字符串列表 words, 返回所有的索引对 [i, j] 使得在索引对范围内的子字符串 text[i]...text[j]（包括 i 和 j）属于字符串列表 words。
*/
func indexPairs(text string, words []string) [][]int {
	var t trie
	for _, v := range words {
		t.Insert(v)
	}

	var result [][]int
	var count int
	var ok bool
	for i := 0; i < len(text); i++ {
		m := make(map[string]bool)
		ok = true
		for ok {
			count, ok = t.exits(text[i:], m)
			if ok {
				result = append(result, []int{i, i + count - 1})
			}
		}
	}
	return result
}
