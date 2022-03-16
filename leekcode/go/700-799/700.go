package _00_799


/*
700. 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
*/

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

    if root == nil {
        return nil
    }

    if root.Val == val {
        return root
    }

    if root.Val > val {
        return searchBST(root.Left, val)
    }

    if root.Val < val {
        return searchBST(root.Right, val)
    }

    return nil
}
