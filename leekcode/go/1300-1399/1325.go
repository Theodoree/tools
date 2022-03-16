package _300_1399



/*
1325. 删除给定值的叶子节点

给你一棵以 root 为根的二叉树和一个整数 target ，请你删除所有值为 target 的 叶子节点 。

注意，一旦删除值为 target 的叶子节点，它的父节点就可能变成叶子节点；如果新叶子节点的值恰好也是 target ，那么这个节点也应该被删除。

也就是说，你需要重复此过程直到不能继续删除。
*/
func removeLeafNodes(root *TreeNode, target int) *TreeNode {

    Remove(root, root, target)
    if root.Left == nil && root.Right == nil && root.Val == target {
        return  nil
    }
    return root
}

func Remove(root *TreeNode, parent *TreeNode, target int) *TreeNode {
    if root == nil {
        return nil
    }

    root.Left = Remove(root.Left, root, target)
    root.Right = Remove(root.Right, root, target)

    if root.Val == target && root.Left == nil && root.Right == nil {
        if parent.Left == root {
            parent.Left = nil
        } else {
            parent.Right = nil
        }
        return nil
    }


    return root
}
