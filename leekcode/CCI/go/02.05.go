package _go


/*
面试题 02.05. 链表求和
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。



示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：假设这些数位是正向存放的，请再做一遍。

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
*/


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var arr []int

    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            arr = append(arr, l1.Val+l2.Val)
            l1 = l1.Next
            l2 = l2.Next
        } else if l1 != nil {
            arr = append(arr, l1.Val)
            l1 = l1.Next
        } else if l2 != nil {
            arr = append(arr, l2.Val)
            l2 = l2.Next
        }
    }

    for i := 0; i < len(arr); i++ {
        if arr[i] >= 10 {
            arr[i] %= 10
            if i == len(arr)-1 {
                arr = append(arr, 1)
            } else {
                arr[i+1]++
            }
        }
    }

    var s = &ListNode{
        Val:  0,
        Next: nil,
    }
    var cur *ListNode

    for i := 0; i < len(arr); i++ {
        if cur == nil {
            s.Next = &ListNode{Val: arr[i],}
            cur = s.Next
        }else{
            cur.Next = &ListNode{Val: arr[i],}
            cur = cur.Next
        }
    }

    return s.Next
}