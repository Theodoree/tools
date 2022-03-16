package _go

/*
剑指 Offer II 025. 链表中的两数相加
给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

可以假设除了数字 0 之外，这两个数字都不会以零开头。



示例1：



输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]
示例2：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]
示例3：

输入：l1 = [0], l2 = [0]
输出：[0]
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr []*ListNode
	var arr1 []*ListNode
	for l1 != nil {
		arr = append(arr, l1)
		l1 = l1.Next
	}

	for l2 != nil {
		arr1 = append(arr1, l2)
		l2 = l2.Next
	}

	if len(arr1) > len(arr) {
		arr, arr1 = arr1, arr
	}
	var plus int
	var result int

	for i := 1; i <= len(arr); i++ {
		result = arr[len(arr)-i].Val
		if len(arr1)-i >= 0 {
			result = arr[len(arr)-i].Val + arr1[len(arr1)-i].Val
		}
		arr[len(arr)-i].Val = result%10 + plus
		plus = 0
		if result >= 10 {
			plus = 1
		}
		if arr[len(arr)-i].Val >= 10 {
			plus++
			arr[len(arr)-i].Val %= 10
		}
	}

	if plus > 0 {
		var l ListNode
		l.Val = plus
		l.Next = arr[0]
		arr = append([]*ListNode{&l}, arr...)
	}

	return arr[0]
}
