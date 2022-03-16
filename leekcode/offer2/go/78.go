package _go

import "math"

/*
剑指 Offer II 078. 合并排序链表
给定一个链表数组，每个链表都已经按升序排列。

请将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}

	var sentry ListNode
	cur := &sentry

	for  {
		min := math.MaxInt64
		var idx int
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[i].Val < min {
				min = lists[i].Val
				idx = i
			}
		}
		if min == math.MaxInt64 {
			break
		}
		cur.Next = lists[idx]
		cur = cur.Next

		lists[idx] = lists[idx].Next
	}

	return sentry.Next

}
