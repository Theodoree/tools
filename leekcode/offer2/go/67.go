package _go

/*
剑指 Offer II 067. 最大的异或
给定一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
示例 1：
输入：nums = [3,10,5,25,2,8]
输出：28
解释：最大运算结果是 5 XOR 25 = 28.
*/
type Trie struct {
	child [2]*Trie
}

func Constructor() *Trie {
	return &Trie{
		child: [2]*Trie{},
	}
}

func (t *Trie) Insert(num int) {
	curr := t
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		if curr.child[bit] == nil {
			curr.child[bit] = Constructor()
		}
		curr = curr.child[bit]
	}
}

func (t *Trie) Check(num int) int {
	var res int
	curr := t
	for i := 30; i >= 0; i-- {
		bit := (num >> i) & 1
		if curr.child[bit^1] != nil {
			res |= 1 << i
			curr = curr.child[bit^1]
		} else {
			curr = curr.child[bit]
		}
	}
	return res
}

func findMaximumXOR(nums []int) int {
	t := Constructor()
	t.Insert(nums[0])
	var res int
	for _, num := range nums[1:] {
		res = max(res, t.Check(num))
		t.Insert(num)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
