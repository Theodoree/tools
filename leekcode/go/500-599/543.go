package _00_599

import (
    "math"
)

/*
543. 二叉树的直径

给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

func diameterOfBinaryTree(root *TreeNode) int {
    var result int
    getDepth(root, &result)
    return result
}

func getDepth(root *TreeNode, result *int) int {

    if root == nil {
        return -1
    }

    left := getDepth(root.Left, result) + 1
    right := getDepth(root.Right, result) + 1
    *result = int(math.Max(float64(left+right), float64(*result)))

    return int(math.Max(float64(left), float64(right)))
}
