package _go

/*
剑指 Offer 06. 从尾到头打印链表

输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]
*/

func reversePrint(head *ListNode) []int {

	var arr []int
	for head != nil {
		arr = append([]int{head.Val}, arr...)
		head = head.Next
	}

	return arr

}
