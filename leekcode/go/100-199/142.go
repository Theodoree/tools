package _00_199

import (
    "fmt"
)

/*
142. 环形链表 II


给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
*/

func detectCycle(head *ListNode) *ListNode {

    v := make(map[string]struct{})

    cur := head
    // 取个巧 内存地址是唯一的
    for cur != nil {
        address := fmt.Sprintf("%p", cur)
        if _, ok := v[address]; ok && address != "" {
            return cur
        } else {
            v[address] = struct{}{}
        }
        cur = cur.Next
    }
    return nil
}

// 双指针
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    slow := head
    fast := head

    for true {
        if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil { // 检测是否有环
            return nil
        }
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast { // 确定有环
            break
        }
    }

    fast = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }

    return fast
}
