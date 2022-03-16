package __99

import "fmt"

//N皇后 II
func totalNQueens(n int) int {

    var array = [][]int{}

    for i := 0; i < n; i++ {
        array = append(array, make([]int, n))
    }

    var times int

    TotalNQueens(0, n, array, &times)

    return times
}

func TotalNQueens(current int, n int, Array [][]int, cnt *int) {

    if current >= n {
        //for _,v:=range Array{
        //fmt.Printf("%v \n",v)
        //}
        //fmt.Println()
        *cnt++
    }

    for i := 0; i < n; i++ {
        if check(current, i, n, Array) {
            Array[current][i] = 1
            TotalNQueens(current+1, n, Array, cnt)
            Array[current][i] = 0
        }

    }

}

func check(currentRow, index, n int, array [][]int) bool {
    for i := 0; i < n; i++ {
        if array[i][index] == 1 { //检查行列冲突
            return false
        }
    }

    m := index - 1
    for i := currentRow - 1; i >= 0 && m >= 0; {
        if array[i][m] == 1 { //检查左对角线
            return false
        }

        i--
        m--
    }

    m = index + 1

    for i := currentRow - 1; i >= 0 && m < n; {
        if array[i][m] == 1 { //检查右对角线
            return false
        }

        i--
        m++
    }

    return true

}

func main() {
    fmt.Println(totalNQueens(4))

}
