package _go


/*
面试题28. 对称的二叉树

请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3



示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false
*/

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return IsSymmetric(root.Left,root.Right)
}

func IsSymmetric(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    if left.Val == right.Val {
        return IsSymmetric(left.Left, right.Right) && IsSymmetric(left.Right, right.Left)
    }

    return false
}
