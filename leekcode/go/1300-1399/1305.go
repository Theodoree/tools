package _300_1399

import "sort"

/*
1305. 两棵二叉搜索树中的所有元素

给你 root1 和 root2 这两棵二叉搜索树。

请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.



示例 1：



输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]
示例 2：

输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
输出：[-10,0,0,1,2,5,7,10]
示例 3：

输入：root1 = [], root2 = [5,1,7,0,2]
输出：[0,1,2,5,7]
示例 4：

输入：root1 = [0,-10,10], root2 = []
输出：[-10,0,10]
*/

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    var arr1, arr2 []int

    DFS(root1, &arr1)
    DFS(root2, &arr2)

    arr1 = append(arr1,arr2...)
    sort.Ints(arr1)
    return arr1
}

func DFS(r *TreeNode, val *[]int) {
    if r == nil {
        return
    }
    DFS(r.Left, val)
    *val = append(*val, r.Val)
    DFS(r.Right, val)
}
