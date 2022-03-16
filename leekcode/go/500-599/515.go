package _00_599


/*
15. 在每个树行中找最大值

您需要在二叉树的每一行中找到最大的值。

示例：

输入:

          1
         / \
        3   2
       / \   \
      5   3   9

输出: [1, 3, 9]
*/
func largestValues(root *TreeNode) []int {

    var result []int

    LargestValues(root, &result, 0)

    return result
}

func LargestValues(root *TreeNode, result *[]int, depth int) {
    if root == nil {
        return
    }

    if len(*result) <= depth {
        *result = append(*result, root.Val)
    } else {
        if t := (*result)[depth]; t < root.Val {
            (*result)[depth] = root.Val
        }
    }

    LargestValues(root.Left, result, depth+1)
    LargestValues(root.Right, result, depth+1)

}
