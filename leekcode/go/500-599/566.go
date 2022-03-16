package _00_599


/*
566. 重塑矩阵

在MATLAB中，有一个非常有用的函数 reshape，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。

给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。

重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。

如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。
*/

func matrixReshape(nums [][]int, r int, c int) [][]int {

    if r*c != len(nums)*len(nums[0]) {
        return nums
    }

    result := make([][]int, 0, r)
    var index int
    l := make([]int, c)

    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums[0]); j++ {
            l[index] = nums[i][j]
            if c-1 == index {
                result = append(result, l)
                l = make([]int, c)
                index = 0
            }else{
                index++
            }
        }
    }
    return result
}