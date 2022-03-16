package _00_199


//112 路径总和

func hasPathSum(root *TreeNode, sum int) bool {
    var result bool
    if root == nil {
        return false
    }
    if root.Val == sum && root.Left == nil && root.Right == nil {
        return true
    }

    HasPathSum(root, sum, 0, &result, 0, false)
    return result

}

func HasPathSum(root *TreeNode, sum int, current int, result *bool, depth int, isHaveuncle bool) {
    if root == nil {

        if current == sum && depth > 1 && !isHaveuncle{
            *result = true
        }
        return
    }
    isHaveuncle = false
    if root.Right != nil {
        isHaveuncle = true
    }

    HasPathSum(root.Left, sum, current+root.Val, result, depth+1, isHaveuncle)
    if root.Left != nil {
        isHaveuncle = true
    }else{
        isHaveuncle = false
    }

    HasPathSum(root.Right, sum, current+root.Val, result, depth+1, isHaveuncle)

}