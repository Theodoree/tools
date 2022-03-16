package _00_799

/*
745. 前缀和后缀搜索
设计一个包含一些单词的特殊词典，并能够通过前缀和后缀来检索单词。

实现 WordFilter 类：

WordFilter(string[] words) 使用词典中的单词 words 初始化对象。
f(string prefix, string suffix) 返回词典中具有前缀 prefix 和后缀suffix 的单词的下标。如果存在不止一个满足要求的下标，返回其中 最大的下标 。如果不存在这样的单词，返回 -1 。


示例

输入：
["WordFilter", "f"]
[[["apple"]], ["a", "e"]]
输出：
[null, 0]

解释：
WordFilter wordFilter = new WordFilter(["apple"]);
wordFilter.f("a", "e"); // 返回 0 ，因为下标为 0 的单词的 prefix = "a" 且 suffix = 'e" 。

提示：

1 <= words.length <= 15000
1 <= words[i].length <= 10
1 <= prefix.length, suffix.length <= 10
words[i]、prefix 和 suffix 仅由小写英文字母组成
最多对函数 f 进行 15000 次调用
通过次数3,012提交次数7,584
请问您在哪类招聘中遇到此题？

*/

type trie struct {
	root node
}

type node struct {
	node [26]*node
	str  string
	idx  int
	end  bool
}

func (t *trie) Insert(s string, idx int) {
	cur := &t.root
	for _, v := range s {
		if cur.node[v-'a'] == nil {
			cur.node[v-'a'] = &node{}
		}
		cur = cur.node[v-'a']
	}
	cur.end = true
	cur.idx = idx
	cur.str = s
}

func (t *trie) exits(s string) bool {
	cur := &t.root
	for i := 0; i < len(s)-1; i++ {
		cur = cur.node[s[i]-'a']
		if cur == nil || !cur.end {
			return false
		}
	}
	return true
}

type WordFilter struct {
	trie trie
}

func Constructor(words []string) WordFilter {
	var w WordFilter
	for idx, v := range words {
		w.trie.Insert(v, idx)
	}
	return w
}

func (w *WordFilter) F(prefix string, suffix string) int {

	// 先确保prefix存在
	pNode := &w.trie.root
	for _, v := range prefix {
		if pNode == nil {
			return -1
		}
		pNode = pNode.node[v-'a']
	}

	return dfs(pNode, suffix)
}

func dfs(p *node, suffix string) int {
	if p == nil {
		return -1
	}
	if p.end &&
		len(p.str)-len(suffix) >= 0 &&
		p.str[len(p.str)-len(suffix):] == suffix {
		return p.idx
	}

	maxIdx := -1
	for _, v := range p.node {
		i := dfs(v, suffix)
		if i > maxIdx {
			maxIdx = i
		}
	}

	return maxIdx

}
