package _00_199

//107 二叉树的层次遍历

//使用map和数组的速度相近,如果想提高执行速度,应该换种思路
func levelOrderBottom(root *TreeNode) [][]int {

    var levelMap = [][]int{}

    LevelOrderBottom(root, &levelMap, 1)
    var val [][]int

    for i:=len(levelMap)-1;i>=0;i--{
        val =append(val,levelMap[i])
    }


    return val
}

func LevelOrderBottom(root *TreeNode, levelMap *[][]int, depth int) {

    if root == nil {
        return
    }

    if len(*levelMap) < depth {
        *levelMap = append(*levelMap,[]int{root.Val})
    }else{

        l:=*levelMap
        l[depth-1] = append(l[depth-1],root.Val)
    }

    if root.Left != nil {
        LevelOrderBottom(root.Left, levelMap, depth+1)
    }
    if root.Right != nil {
        LevelOrderBottom(root.Right, levelMap, depth+1)
    }

}
