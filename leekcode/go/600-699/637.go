package _00_699


/*
637. 二叉树的层平均值

给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.

示例 1:

输入:
    3
   / \
  9  20
    /  \
   15   7
输出: [3, 14.5, 11]
解释:
第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
*/


func averageOfLevels(root *TreeNode) []float64 {
    var result [][]float64
    AverageOfLevels(root, &result, 0)

    var r []float64
    for i := 0; i < len(result); i++ {
        r = append(r, float64(result[i][0])/float64(result[i][1]))
    }

    return r
}
func AverageOfLevels(root *TreeNode, result *[][]float64, depth int) {
    if root == nil {
        return
    }
    if len(*result) <= depth {
        *result = append(*result, make([]float64, 2, 2))
    }
    current := *result
    current[depth][0] += float64(root.Val)
    current[depth][1] += 1

    AverageOfLevels(root.Left, result, depth+1)
    AverageOfLevels(root.Right, result, depth+1)

}