package _00_299


/*
240. 搜索二维矩阵 II
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。


*/


func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    var (
        width  = len(matrix[0]) - 1
        height = len(matrix) - 1
        h      int
    )
    for width >= 0 && h <= height {
        if matrix[h][width] > target {
            width--
        } else if matrix[h][width] < target {
            h++
        } else {
            return true
        }
    }
    return false
}