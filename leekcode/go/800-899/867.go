package _00_899



/*
867. 转置矩阵
给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

 

示例 1：

输入：[[1,2,3],
      [4,5,6],
      [7,8,9]]
输出：[[1,4,7],
      [2,5,8],
      [3,6,9]]
示例 2：

输入：[[1,2,3],
      [4,5,6]]
输出：[[1,4],
      [2,5],
      [3,6]]

*/

func transpose(A [][]int) [][]int {
    var results [][]int

    for i := 0; i < len(A[0]); i++ {
        var result []int
        index := 0
        for index < len(A) {
            result = append(result, A[index][i])
            index++
        }
        results = append(results, result)
    }

    return results
}


func transpose1(A [][]int) [][]int {
    if A == nil || len(A) == 0 || len(A[0]) == 0 {
        return [][]int{}
    }
    m := len(A)
    n := len(A[0])
    newA := make([][]int, n)
    for i := 0; i < n; i++ {
        newA[i] = make([]int, m)
        for j := 0; j < m; j++ {
            newA[i][j] = A[j][i]
        }
    }
    return newA
}
