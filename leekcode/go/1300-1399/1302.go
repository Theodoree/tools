package _300_1399


/* 1302. 层数最深叶子节点的和
给你一棵二叉树，请你返回层数最深的叶子节点的和。

*/

func deepestLeavesSum(root *TreeNode) int {
    var arr [][]int
    dfs(&arr, 0, root)

    var sum int
    for i := 0; i < len(arr[len(arr)-1]); i++ {
        sum += arr[len(arr)-1][i]
    }
    return sum
}

func dfs(arr *[][]int, depth int, root *TreeNode) {
    if root == nil {
        return
    }
    if len(*arr) <= depth {
        *arr = append(*arr, []int{root.Val})
    } else {
        cur := *arr
        cur[depth] = append(cur[depth], root.Val)
    }

    dfs(arr, depth+1, root.Left)
    dfs(arr, depth+1, root.Right)
}
