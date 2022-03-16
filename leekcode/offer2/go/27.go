package _go

import "sync"

/*
剑指 Offer II 027. 回文链表
给定一个链表的 头节点 head ，请判断其是否为回文链表。

如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。



示例 1：



输入: head = [1,2,3,3,2,1]
输出: true
示例 2：



输入: head = [1,2]
输出: false

*/

var p = sync.Pool{New: func() interface{} {
	return make([]int, 0, 10)
}}

func isPalindrome(head *ListNode) bool {
	arr := p.Get().([]int)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	if len(arr) <= 1 {
		return true
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}

	arr = arr[:0]
	p.Put(arr)
	return true

}
