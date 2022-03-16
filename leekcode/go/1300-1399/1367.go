package _300_1399

/*
1367. 二叉树中的列表
给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
*/

func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }

    if head.Val == root.Val {
        if IsSubPath(head, root) {
            return true
        }
    }

    return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func IsSubPath(s *ListNode, t *TreeNode) bool {
    if s == nil {
        return true
    } else if t == nil {
        return false
    } else if s.Val != t.Val {
        return false
    }

    return IsSubPath(s.Next, t.Left) || IsSubPath(s.Next, t.Right)

}
