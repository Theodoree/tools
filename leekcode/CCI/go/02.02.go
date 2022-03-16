package _go

/*
面试题 02.02. 返回倒数第 k 个节点
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
*/

func kthToLast(head *ListNode, k int) int {

    var arr []*ListNode

    for head != nil {
        arr = append(arr, head)
        head = head.Next
    }

    return  arr[len(arr)-k].Val
}
