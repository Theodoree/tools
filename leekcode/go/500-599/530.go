package _00_599


/*
530. 二叉搜索树的最小绝对差

给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

示例 :

输入:

   1
    \
     3
    /
   2

输出:
1

解释:
最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
*/

func getMinimumDifference(root *TreeNode) int {

    var v []int
    getValues(root, &v)
    var min int
    min = math.MaxInt32
    for i := 1; i < len(v); i++ {

        tmp := v[i] - v[i-1]
        if tmp < min {
            min = tmp
        }
    }

    if min == math.MaxInt32 {
        return 0
    }

    return min

}
func getValues(root *TreeNode, result *[]int) {
    if root == nil {
        return
    }
    // 这里选择中序遍历是因为该树是搜索树
    getValues(root.Left, result)
    *result = append(*result, root.Val)
    getValues(root.Right, result)

}
