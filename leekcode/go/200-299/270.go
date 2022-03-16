package _00_299

import "math"

func closestValue(root *TreeNode, target float64) int {

    if root == nil {
        return 0
    }

    var result []int
    GetSlice(root, &result)

    var current int
    if len(result) > 0 {
        current = result[0]
    }
    for i := 1; i < len(result); i++ {
        if math.Abs(float64(current)-target) > math.Abs(float64(result[i])-target) {
            current = result[i]
        }

    }

    return current
}

func GetSlice(root *TreeNode, f *[]int) {

    if root == nil {
        return
    }

    *f = append(*f, root.Val)

    GetSlice(root.Left, f)
    GetSlice(root.Right, f)

}

