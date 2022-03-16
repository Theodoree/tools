package _00_999

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

/*
938. 二叉搜索树的范围和

给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。

二叉搜索树保证具有唯一的值。

示例 1：

输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
输出：32
示例 2：

输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
输出：23
*/
func rangeSumBST(root *TreeNode, L int, R int) int {
    if root == nil {
        return 0
    }
    if root.Val > R {
        return rangeSumBST(root.Left, L, R)
    } else if root.Val < L {
        return rangeSumBST(root.Right, L, R)
    } else {
        return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
    }
}

