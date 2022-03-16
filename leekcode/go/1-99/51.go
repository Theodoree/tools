package __99

import (
    "fmt"
    "strings"
)

//51.N皇后
func solveNQueens(n int) [][]string {

    var array = [][]string{}

    for i := 0; i < n; i++ {
        var t []string
        for j := 0; j < n; j++ {
            t = append(t, `.`)
        }
        array = append(array, t)
    }

    var ans [][]string
    SolveNQueens(0, n, array, &ans)

    return ans

}

func SolveNQueens(current, n int, Array [][]string, ans *[][]string) {
    if current >= n {
        var str []string
        for _,v:=range Array{
            str = append(str,strings.Join(v,``))
        }

        *ans = append(*ans, str)
    }

    for i := 0; i < n; i++ {
        if check1(current, i, n, Array) {
            Array[current][i] = `Q`
            SolveNQueens(current+1, n, Array, ans)
            Array[current][i] = `.`
        }

    }

}

func check1(currentRow, index, n int, array [][]string) bool {
    for i := 0; i < n; i++ {
        if array[i][index] == "Q" { //检查行列冲突
            return false
        }
    }

    m := index - 1
    for i := currentRow - 1; i >= 0 && m >= 0; {
        if array[i][m] == "Q" { //检查左对角线
            return false
        }

        i--
        m--
    }

    m = index + 1

    for i := currentRow - 1; i >= 0 && m < n; {
        if array[i][m] == "Q" { //检查右对角线
            return false
        }

        i--
        m++
    }

    return true

}


func main(){
    fmt.Println(solveNQueens(4))


}
