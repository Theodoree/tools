package _00_699

/*
676. 实现一个魔法字典
设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。
实现 MagicDictionary 类：
MagicDictionary() 初始化对象
void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字符串互不相同
bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。
*/
type MagicDictionary struct {
	root node1
}

type node1 struct {
	node [26]*node1
	cnt  int
	end  bool
}

func (t *node1) Insert(s string) {
	cur := t
	for _, v := range s {
		if cur.node[v-'a'] == nil {
			cur.cnt++
			cur.node[v-'a'] = &node1{}
		}
		cur = cur.node[v-'a']
	}
	cur.end = true
}

func (t *node1) exits(s string, flag bool) bool {
	cur := t

	for i := 0; i < len(s); i++ {
		if cur == nil {
			return false
		}
		parent := cur
		cur = cur.node[s[i]-'a']

		if !flag {
			for _, v := range parent.node {
				if v != cur && v != nil && v.exits(s[i+1:], true) {
					return true
				}
			}
		}

	}
	if cur == nil {
		return false
	}
	return cur.end && flag
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}

}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for i := 0; i < len(dictionary); i++ {
		this.root.Insert(dictionary[i])
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return this.root.exits(searchWord, false)
}



