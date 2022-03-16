package _00_699

import (
    "sort"
)

/*
671. 二叉树中第二小的节点

给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。

给出这样的一个二叉树，你需要输出所有节点中的第二小的值。如果第二小的值不存在的话，输出 -1 。

示例 1:

输入:
    2
   / \
  2   5
     / \
    5   7

输出: 5
说明: 最小的值是 2 ，第二小的值是 5 。
示例 2:

输入:
    2
   / \
  2   2

输出: -1
说明: 最小的值是 2, 但是不存在第二小的值。
*/

func findSecondMinimumValue(root *TreeNode) int {

    var v = make(map[int]struct{})
    FindSecondMinimumValue(root, v)

    var result []int
    for k, _ := range v {
        result = append(result, k)
    }

    sort.Slice(result, func(i, j int) bool {
        return result[i] < result[j]
    })

    if len(result) < 2 {
        return -1
    }

    return result[1]

}

func FindSecondMinimumValue(root *TreeNode, result map[int]struct{}) {
    if root == nil {
        return
    }

    result[root.Val] = struct{}{}
    FindSecondMinimumValue(root.Left, result)
    FindSecondMinimumValue(root.Right, result)

}
