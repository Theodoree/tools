package _00_199

//145. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {

    var result []int
    PreorderTraversal(root, &result)

    return result
}

func PreorderTraversal(t *TreeNode, result *[]int) {
    if t == nil {
        return
    }

    *result = append(*result, t.Val)
    PreorderTraversal(t.Left, result)
    PreorderTraversal(t.Right, result)
}