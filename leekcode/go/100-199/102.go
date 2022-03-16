package _00_199


//102. 二叉树的层次遍历
func levelOrder(root *TreeNode) [][]int {
    var slice [][]int
    LevelOrder(root, 0, &slice)
    return slice
}

func LevelOrder(root *TreeNode, depth int, slice *[][]int) {
    if root == nil {
        return
    }

    if len( *slice) < depth+1 {
        *slice = append( *slice, []int{root.Val})
    }else{
        array :=*slice
        array[depth] = append(array[depth], root.Val)
    }


    LevelOrder(root.Left, depth+1, slice)
    LevelOrder(root.Right, depth+1, slice)

}
