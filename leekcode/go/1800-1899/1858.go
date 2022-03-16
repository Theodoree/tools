package _800_1899



/*
1858. 包含所有前缀的最长单词
给定一个字符串数组 words，找出 words 中所有的前缀都在 words 中的最长字符串。
例如，令 words = ["a", "app", "ap"]。字符串 "app" 含前缀 "ap" 和 "a" ，都在 words 中。
返回符合上述要求的字符串。如果存在多个（符合条件的）相同长度的字符串，返回字典序中最小的字符串，如果这样的字符串不存在，返回 ""。
*/

type trie struct {
	child [26]*trie
	end   bool
}

func (t *trie) Insert(s string) bool {
	flag:=true
	cur := t
	for idx, v := range s {
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &trie{}
		}
		cur = cur.child[v-'a']
		if idx !=len(s)-1 && !cur.end{
			flag = false
		}
	}
	cur.end = true
	return flag
}

func longestWord(words []string) string {
	sort.Strings(words)
	var t trie
	var str string
	for _,v:=range words{
		if t.Insert(v) {
			if len(v) > len(str){
				str = v
			}else if len(v) == len(str) && str > v {
				v = str
			}
		}
	}
	return str
}

