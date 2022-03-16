package _00_199

/*
105. 从前序与中序遍历序列构造二叉树

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7


前序 DLR   第一个输出的节点               在中序中,左边是左子树,右边是右子树
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
    if leftIdx == RightIdx  {
        return nil
    }


    root := &TreeNode{Val: preorder[*cnt]}
    idx := m[root.Val]
    *cnt++
    root.Left = BuildTree(preorder, inorder, cnt, m, leftIdx, idx)
    root.Right = BuildTree(preorder, inorder, cnt, m, idx+1, RightIdx)

    return root

}

