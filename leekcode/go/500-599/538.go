package _00_599


/*
538. 把二叉搜索树转换为累加树

给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

例如：

输入: 二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
    sum := 0
    f(root, sum)
    return root
}

func f(root *TreeNode, sum int) int {
    if root == nil {
        return sum
    }
    sum = f(root.Right, sum)
    root.Val += sum
    sum = root.Val
    sum = f(root.Left, sum)
    return sum
}

