package _000_1099

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

func queryString(S string, N int) bool {

    switch {
    case N < 0:
        return false
    case len(S) == 0 && N != 0:
        return false
    case len(S) == 0 && N == 0:
        return true
    }
    var sum int

    strS := strings.Split(S, ``)
    cnt := 0
    for i := len(strS) - 1; i >= 0; i-- {
        v, err := strconv.Atoi(strS[i])
        if err != nil {
            fmt.Println(err)
        }
        sum+= int(math.Pow(2,float64(cnt)))*v
        cnt++
    }

    return sum == N
}

func main() {
    fmt.Println(queryString("0011", 3))

}


