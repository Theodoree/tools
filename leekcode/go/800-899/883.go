package _00_899



/*
883. 三维形体投影面积

在 N * N 的网格中，我们放置了一些与 x，y，z 三轴对齐的 1 * 1 * 1 立方体。

每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。

现在，我们查看这些立方体在 xy、yz 和 zx 平面上的投影。

投影就像影子，将三维形体映射到一个二维平面上。

在这里，从顶部、前面和侧面看立方体时，我们会看到“影子”。

返回所有三个投影的总面积。



示例 1：

输入：[[2]]
输出：5
示例 2：

输入：[[1,2],[3,4]]
输出：17
解释：
这里有该形体在三个轴对齐平面上的三个投影(“阴影部分”)。
*/

func projectionArea(grid [][]int) int {

    var sum int

    for i := 0; i < len(grid); i++ {
        var row, col int

        for j := 0; j < len(grid); j++ {
            if grid[i][j] > 0 {
                sum++
            }

            if grid[i][j] > row {
                row = grid[i][j]
            }

            if grid[j][i] > col {
                col = grid[j][i]
            }
        }

        sum += row + col
    }

    return  sum
}
