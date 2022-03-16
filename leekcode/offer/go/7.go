package _go

/*
剑指 Offer 07. 重建二叉树

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。



例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7



前序 DLR   第一个输出的节点肯定是root
中序 LDR   第一个输出的肯定是left的叶子节点
后续 LRD   第一个输出的肯定是Right的右节点
*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	var cnt int
	var m = make(map[int]int)

	for k, v := range inorder {

		m[v] = k
	}

	return BuildTree(preorder, inorder, &cnt, m, 0, len(inorder))
}

func BuildTree(preorder []int, inorder []int, cnt *int, m map[int]int, leftIdx, RightIdx int) *TreeNode {
	if leftIdx == RightIdx {
		return nil
	}

	root := &TreeNode{Val: preorder[*cnt]}
	idx := m[root.Val]
	*cnt++
	root.Left = BuildTree(preorder, inorder, cnt, m, leftIdx, idx)
	root.Right = BuildTree(preorder, inorder, cnt, m, idx+1, RightIdx)

	return root

}
