package _go

/*
面试题54. 二叉搜索树的第k大节点

给定一棵二叉搜索树，请找出其中第k大的节点。

示例 1:

输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:

输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
*/

func kthLargest(root *TreeNode, k int) int {
    var arr []int
    DFS(root,&arr,k)
    return  arr[k]
}

func DFS(root *TreeNode, cur *[]int,k int)  {
    if root == nil || len(*cur) == k {
        return
    }
    DFS(root.Right,cur, k)
    *cur = append(*cur,root.Val)
    DFS(root.Left, cur,k)
}
