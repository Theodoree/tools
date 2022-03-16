package __99

import (
    "fmt"
)

/*
59. 螺旋矩阵 II

给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {

    var Matrix [][]int

    Matrix = make([][]int, n)

    for i := 0; i < n; i++ {
        Matrix[i] = make([]int, n)
    }

    var flag int
    var row int
    var l int
    flag = 1
    // 1 向左 2向下 3向右 4 向上

    for i := 1; i <= n*n; i++ {

        switch flag {
        case 1:
            Matrix[row][l] = i
            if l+1 == n || Matrix[row][l+1] != 0 {
                row++
                flag = 2
            } else {
                l++
            }
        case 2:
            Matrix[row][l] = i

            if row+1 == n || Matrix[row+1][l] != 0 {
                l--
                flag = 3
            } else {
                row++
            }
        case 3:
            Matrix[row][l] = i
            if l <= 0 || Matrix[row][l-1] != 0 {
                row--
                flag = 4
            } else {
                l--
            }
        case 4:
            Matrix[row][l] = i
            if row <= 0 || Matrix[row-1][l] != 0 {
                l++
                flag = 1
            } else {
                row--
            }
        }

    }

    return Matrix
}

func main() {
    for _, v := range generateMatrix(10) {
        fmt.Println(v)
    }
}