package _00_499



/*
445. 两数相加 II

给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。



你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:

如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:

输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    var l1Arr []int
    var l2Arr []int

    for l1 != nil {
        l1Arr = append(l1Arr, l1.Val)
        l1 = l1.Next
    }

    for l2 != nil {
        l2Arr = append(l2Arr, l2.Val)
        l2 = l2.Next
    }

    for len(l1Arr) > len(l2Arr) {
        l2Arr = append([]int{0}, l2Arr...)
    }

    for len(l2Arr) > len(l1Arr) {
        l1Arr = append([]int{0}, l1Arr...)
    }

    var result []int
    var cnt int
    for i := len(l1Arr) - 1; i >= 0; i-- {

        sum := l1Arr[i] + l2Arr[i] + cnt
        cnt = 0
        cnt = sum / 10
        result = append(result, sum%10)
    }

    if cnt > 0 {
        result = append(result, cnt)
    }

    var head *ListNode
    if len(result) > 0 {
        head = &ListNode{Val: result[len(result)-1]}
        cur := head

        for i := len(result) - 2; i >= 0; i-- {
            cur.Next = &ListNode{Val: result[i]}
            cur = cur.Next
        }

    }

    return head
}
// 上面跟下面都是一样O(n)的时间复杂度,上面的空间复杂度是O(n) 下面是 原地算法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }

    if l2 == nil {
        return l1
    }

    len1 := 0
    c1 := l1
    len2 := 0
    c2 := l2

    for c1 != nil {
        len1 ++
        c1 = c1.Next
    }

    for c2 != nil {
        len2++
        c2=c2.Next
    }

    if len1 < len2 {
        len1,len2 = len2, len1
        l1,l2 = l2,l1
    }

    c1 = l1
    for i:= 0; i < len1-len2; i++ {
        c1 = c1.Next
    }

    c2 = l2
    for c1 != nil {
        c1.Val += c2.Val
        c1 = c1.Next
        c2 = c2.Next
    }

    flag := true
    for flag {
        flag = false
        c1 = l1
        for c1 != nil && c1.Next != nil {
            c1.Val += c1.Next.Val / 10
            if c1.Next.Val > 9 {
                flag = true
            }
            c1.Next.Val %= 10
            c1 = c1.Next
        }
    }

    if l1.Val >= 10 {
        ret := ListNode{l1.Val/10, l1}
        l1.Val %= 10
        ret.Next = l1
        return &ret
    }else{
        return l1
    }
}
