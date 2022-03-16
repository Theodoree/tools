package _00_799

/*
766. 托普利茨矩阵

如果一个矩阵的每一方向由左上到右下的对角线上具有相同元素，那么这个矩阵是托普利茨矩阵。

给定一个 M x N 的矩阵，当且仅当它是托普利茨矩阵时返回 True。
输入:
matrix = [
  [1,2,3,4],
  [5,1,2,3],
  [9,5,1,2]
]
输出: True
解释:
在上述矩阵中, 其对角线为:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]"。
各条对角线上的所有元素均相同, 因此答案是True。

*/


func isToeplitzMatrix(matrix [][]int) bool {

    if len(matrix) == 0 || len(matrix) == 1 || len(matrix[0]) == 1 {
        return true
    }

    row := len(matrix) - 2
    times := len(matrix) + len(matrix[0]) - 2
    var flag bool
    var r, i, index int

    for times > 0 {
        r = row
        i = index
        for r < len(matrix)-1 && i < len(matrix[0])-1 {
            if matrix[r][i] != matrix[r+1][i+1] {
                return false
            }
            r++
            i++
        }

        if !flag {
            row--
        } else {
            index++
        }

        times--

        if row <= 0 && !flag {
            flag = true
            row = 0
        }

    }

    return true
}
