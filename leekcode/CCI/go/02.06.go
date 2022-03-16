package _go


/*
面试题 02.06. 回文链表
编写一个函数，检查输入的链表是否是回文的。

示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true
*/

func isPalindrome(head *ListNode) bool {
    var arr []int

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    if len(arr) == 1 {
        return true
    }

    if len(arr)%2 == 0 {
        for i := 0; i < len(arr)/2; i++ {
            if arr[i] != arr[len(arr)-1-i] {
                return false
            }
        }
    } else {
        left := len(arr)/2 - 1
        right := len(arr)/2 + 1
        for {

            if arr[left] != arr[right] {
                return false
            }
            left--
            right++

            if left < 0 {
                break
            }
        }

    }

    return true

}