package _00_699


/*
687. 最长同值路径

给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:

2
示例 2:

输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:

2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
*/

func longestUnivaluePath(root *TreeNode) int {

    if root == nil{
        return  0
    }

    var max int
    LongestUnivaluePath(root, &max, root.Val)
    return max

}

func LongestUnivaluePath(root *TreeNode, max *int, f int) int {
    if root == nil {
        return 0
    }

    left := LongestUnivaluePath(root.Left, max, root.Val)
    right := LongestUnivaluePath(root.Right, max, root.Val)
    *max = GetMax(*max, left+right)
    if root.Val == f {
        return GetMax(left, right) + 1
    }
    return 0
}
func GetMax(nums ...int) int {
    var Max int
    Max = nums[0]

    for i := 1; i < len(nums); i++ {
        if nums[i] > Max {
            Max = nums[i]
        }

    }
    return Max
}
