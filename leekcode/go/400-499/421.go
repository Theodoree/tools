package _00_499

/*
421. 数组中两个数的最大异或值
给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
进阶：你可以在 O(n) 的时间解决这个问题吗？
*/
type trie struct {
	child [2]*trie
}
func (t *trie) Insert(cnt int) {
	root:=t
	for i:=63;i>=0;i--{
		idx:= 0
		if cnt&(1 << i) > 0 {
			idx =1
		}
		if root.child[idx] == nil {
			root.child[idx] = &trie{}
		}
		root = root.child[idx]
	}
}

func (t *trie) getMax(cnt int) int{
	root:=t

	for i:=63;i>=0;i--{
		idx:= 1
		if cnt&(1 << i) > 0 {
			idx =0
		}
		if root.child[idx] == nil {
			idx = idx^1
		}
		root = root.child[idx]
		if idx == 1 {
			cnt^= 1 <<i
		}
	}

	return cnt
}

func findMaximumXOR(nums []int) int {

	var t trie
	var max int

	for _,v:=range nums{
		t.Insert(v)
		m:=t.getMax(v)
		if m > max {
			max = m
		}
	}

	return max
}
