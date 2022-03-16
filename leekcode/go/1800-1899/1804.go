package _800_1899


/*

1804. 实现 Trie （前缀树） II
前缀树（trie ，发音为 "try"）是一个树状的数据结构，用于高效地存储和检索一系列字符串的前缀。前缀树有许多应用，如自动补全和拼写检查。

实现前缀树 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 将字符串 word 插入前缀树中。
int countWordsEqualTo(String word) 返回前缀树中字符串 word 的实例个数。
int countWordsStartingWith(String prefix) 返回前缀树中以 prefix 为前缀的字符串个数。
void erase(String word) 从前缀树中移除字符串 word 。
*/

type Trie struct {
	r *root
}

type root struct{
	child [26]*root
	cnt   	 int
	endCnt   int
}

func Constructor() Trie {
	return Trie{r: &root{}}
}


func (t *Trie) Insert(word string)  {
	node:=t.r
	for i:=0;i<len(word);i++{
		if node.child[word[i]-'a'] == nil {
			node.child[word[i]-'a'] = &root{}
		}
		node = node.child[word[i]-'a']
		node.cnt++
	}
	node.endCnt++
}


func (t *Trie) CountWordsEqualTo(word string) int {
	node:=t.r
	for i:=0;i<len(word);i++{
		node = node.child[word[i]-'a']
		if node == nil {
			return 0
		}
	}
	return node.endCnt
}


func (t *Trie) CountWordsStartingWith(prefix string) int {
	node:=t.r
	for i:=0;i<len(prefix);i++{
		node = node.child[prefix[i]-'a']
		if node == nil {
			return 0
		}
	}
	return node.cnt
}


func (t *Trie) Erase(word string)  {
	node:=t.r
	for i:=0;i<len(word);i++{
		node = node.child[word[i]-'a']
		node.cnt--
	}
	node.endCnt--
}

