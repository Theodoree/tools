package _go


/*
面试题27. 二叉树的镜像

请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
镜像输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1



示例 1：

输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
*/

func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    MirrorTree(root.Left, root.Right, root)
    return root

}

func MirrorTree(left, right *TreeNode, parent *TreeNode) {
    if left == nil && right == nil {
        return
    }
    parent.Left, parent.Right = parent.Right, parent.Left
    left = parent.Left
    right = parent.Right
    if left != nil {
        MirrorTree(left.Left, left.Right, left)
    }
    if right != nil {
        MirrorTree(right.Left, right.Right, right)
    }
}
