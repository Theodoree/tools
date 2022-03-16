package _00_199


//145. 二叉树的后序遍历

func postorderTraversal(root *TreeNode) []int {
    var result []int
    PostorderTraversal(root,&result)
    return result
}

func PostorderTraversal(t *TreeNode, result *[]int) {
    if t == nil {
        return
    }

    PostorderTraversal(t.Left, result)
    PostorderTraversal(t.Right, result)
    *result = append(*result, t.Val)

}