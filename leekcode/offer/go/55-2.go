package _go

/*
面试题55 - II. 平衡二叉树

输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。
*/
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return  true
    }
    var l, r int
    GetDepth(root.Left, 0, &l)
    GetDepth(root.Right, 0, &r)

    if math.Abs(float64(l-r)) > 1 {
        return  false
    }

    return isBalanced(root.Left) && isBalanced(root.Right)
}

func GetDepth(root *TreeNode, depth int, maxDepth *int) {
    if root == nil {
        if depth > *maxDepth {
            *maxDepth = depth
        }
        return
    }

    GetDepth(root.Left, depth+1, maxDepth)
    GetDepth(root.Right, depth+1, maxDepth)
}

