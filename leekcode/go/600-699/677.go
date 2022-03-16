package _00_699

/*
677. 键值映射
实现一个 MapSum 类，支持两个方法，insert 和 sum：

MapSum() 初始化 MapSum 对象
void insert(String key, int val) 插入 key-val 键值对，字符串表示键 key ，整数表示值 val 。如果键 key 已经存在，那么原来的键值对将被替代成新的键值对。
int sum(string prefix) 返回所有以该前缀 prefix 开头的键 key 的值的总和。


示例：

输入：
["MapSum", "insert", "sum", "insert", "sum"]
[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
输出：
[null, null, 3, null, 5]

解释：
MapSum mapSum = new MapSum();
mapSum.insert("apple", 3);
mapSum.sum("ap");           // return 3 (apple = 3)
mapSum.insert("app", 2);
mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)


提示：
1 <= key.length, prefix.length <= 50
key 和 prefix 仅由小写英文字母组成
1 <= val <= 1000
最多调用 50 次 insert 和 sum

*/

// 前缀树
type MapSum struct {
	root tree
}

type tree struct {
	child [26]*tree
	val   int
}

func (t *tree) Insert(key string, val int) {
	cur := t
	for _, v := range key {
		if cur.child[v-'a'] == nil {
			cur.child[v-'a'] = &tree{}
		}
		cur = cur.child[v-'a']

	}
	cur.val = val
}

func (t *tree) Sum() int {
	i := t.val
	for _, v := range t.child {
		if v != nil {
			i += v.Sum()
		}
	}
	return i
}

func (t *tree) preSum(prefix string) int {
	cur := t
	for _, v := range prefix {
		if cur == nil {
			return 0
		}
		cur = cur.child[v-'a']
	}

	if cur == nil {
		return 0
	}

	return cur.Sum()
}

func Constructor() MapSum {
	return MapSum{root: tree{}}
}

func (m *MapSum) Insert(key string, val int) {
	m.root.Insert(key, val)
}

func (m *MapSum) Sum(prefix string) int {
	return m.root.preSum(prefix)
}
