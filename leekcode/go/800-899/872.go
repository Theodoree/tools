package _00_899


/*
872. 叶子相似的树

请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。



举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。

如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。

如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。



提示：

给定的两颗树可能会有 1 到 100 个结点。
*/
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

    var result1, result2 []int
    getArray(root1, &result1)
    getArray(root2, &result2)

    if len(result1) != len(result2) {
        return false
    }

    for i:=0;i<len(result1);i++{
        if result1[i] != result2[i]{
            return  false
        }
    }
    return true
}

func getArray(root *TreeNode, result *[]int) {
    if root == nil{
        return
    }
    if root.Left == nil && root.Right == nil {
        *result = append(*result, root.Val)
        return
    }

    getArray(root.Left, result)
    getArray(root.Right, result)


}
