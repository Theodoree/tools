package _00_199


/*
124. 二叉树中的最大路径和

给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。

示例 1:

输入: [1,2,3]

       1
      / \
     2   3

输出: 6
示例 2:

输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42
*/

func maxPathSum(root *TreeNode) int {
    var max int
    max = math.MinInt32
    DFS(root, &max)
    return max
}

func DFS(root *TreeNode, max *int) int {
    if root == nil {
        return 0
    }
    left := DFS(root.Left, max)
    right := DFS(root.Right, max)
    lmr := root.Val + GetMax(0, left) + GetMax(0, right)
    ret := root.Val + GetMax(0, GetMax(left, right))
    *max = GetMax(*max, GetMax(lmr, ret))

    return ret
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

