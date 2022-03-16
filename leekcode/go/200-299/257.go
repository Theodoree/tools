package _00_299

import (
    "fmt"
    "strings"
)

/*
257. 二叉树的所有路径

给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
*/

func binaryTreePaths(root *TreeNode) []string {

    var result []string
    BinaryTreePaths(root, &result, []string{})
    return result

}

func BinaryTreePaths(root *TreeNode, Result *[]string, current []string) {
    if root == nil{
        return
    }
    current = append(current, fmt.Sprintf("%d", root.Val))

    if root.Left == nil && root.Right == nil {
        *Result = append(*Result, strings.Join(current, `->`))
        return
    }

    BinaryTreePaths(root.Left, Result, current)
    BinaryTreePaths(root.Right, Result, current)

}
