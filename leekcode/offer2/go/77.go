package _go

import "sort"

/*
剑指 Offer II 077. 链表排序
给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。



示例 1：



输入：head = [4,2,1,3]
输出：[1,2,3,4]
示例 2：



输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var list []*ListNode
	for head != nil {
		list = append(list,head)
		head = head.Next
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})


	for i:=0;i<len(list)-1;i++{
		list[i].Next = list[i+1]
	}

	list[len(list)-1].Next = nil

	return list[0]




}