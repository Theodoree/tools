package _00_999


/*
993. 二叉树的堂兄弟节点

在二叉树中，根节点位于深度 0 处，每个深度为 k 的节点的子节点位于深度 k+1 处。

如果二叉树的两个节点深度相同，但父节点不同，则它们是一对堂兄弟节点。

我们给出了具有唯一值的二叉树的根节点 root，以及树中两个不同节点的值 x 和 y。

只有与值 x 和 y 对应的节点是堂兄弟节点时，才返回 true。否则，返回 false。



示例 1：


输入：root = [1,2,3,4], x = 4, y = 3
输出：false
示例 2：


输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
输出：true
示例 3：



输入：root = [1,2,3,null,4], x = 2, y = 3
输出：false
*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {

    p, depth := IsCousins(root, root, x, 0)
    p1, depth1 := IsCousins(root, root, y, 0)

    if p == p1 {
        return false
    }

    return depth == depth1

}

func IsCousins(root *TreeNode, parent *TreeNode, x int, depth int) (*TreeNode, int) {
    if root == nil {
        return parent, -1
    }

    if root.Val == x {
        return parent, depth
    }

    p, k := IsCousins(root.Left, root, x, depth+1)
    p1, k1 := IsCousins(root.Right, root, x, depth+1)

    if k > k1 {
        return p, k
    } else {
        return p1, k1
    }
}
