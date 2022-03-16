package _00_399



/*
364. 加权嵌套序列和 II
给你一个整数嵌套列表 nestedList ，每一个元素要么是一个整数，要么是一个列表（这个列表中的每个元素也同样是整数或列表）。
整数的 深度 取决于它位于多少个列表内部。例如，嵌套列表 [1,[2,2],[[3],2],1] 的每个整数的值都等于它的 深度 。令 maxDepth 是任意整数的 最大深度 。
整数的 权重 为 maxDepth - (整数的深度) + 1 。
将 nestedList 列表中每个整数先乘权重再求和，返回该加权和。
*/

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {}
func depthSumInverse(nestedList []*NestedInteger) int {
	var arr [][]int
	var tmp []*NestedInteger
	var cnt int
	for len(nestedList) > 0 {
		if len(arr) <= cnt {
			arr = append(arr, []int{})
		}
		for _, v := range nestedList {
			if v.IsInteger() {
				arr[cnt] = append(arr[cnt], v.GetInteger())
			} else {
				tmp = append(tmp, v.GetList()...)
			}
		}
		nestedList, tmp = tmp, nestedList
		tmp = tmp[:0]
		cnt++
	}

	var sum int

	for idx,subArr:=range arr{
		for _,v:=range subArr{
			sum+=v*(len(arr)-idx)
		}
	}
	return sum

}
