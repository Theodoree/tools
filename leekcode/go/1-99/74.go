package __99


/*
74. 搜索二维矩阵

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    n := len(matrix[0]) - 1
    var i, left int
    var flag bool
    for i = 0; i < len(matrix); i++ {
        if matrix[i][0] <= target && matrix[i][n] >= target {
            flag = true
            break
        }
    }

    if !flag {
        return false
    }

    for left <= n {
        if matrix[i][left] == target || matrix[i][n] == target {
            return true
        }
        left++
        n--
    }

    return false

}

