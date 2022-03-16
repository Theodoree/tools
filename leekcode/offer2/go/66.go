package _go

type trie struct {
	root node
}

type node struct {
	node  [26]*node
	count int
	end   bool
}

func (t *trie) Insert(s string, count int) {
	cur := &t.root
	for _, v := range s {
		if cur.node[v-'a'] == nil {
			cur.node[v-'a'] = &node{}
		}
		cur = cur.node[v-'a']
	}
	cur.end = true
	cur.count = count
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

type MapSum struct {
	t trie
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (m *MapSum) Insert(key string, val int) {
	m.t.Insert(key, val)
}

func (m *MapSum) Sum(prefix string) int {
	cur := &m.t.root
	for _, v := range prefix {
		if cur == nil {
			return 0
		}
		cur = cur.node[v-'a']
	}

	return dfs(cur)
}

func dfs(n *node) int {
	if n == nil {
		return 0
	}
	sum := n.count
	for _, v := range n.node {
		sum += dfs(v)
	}
	return sum
}
