package _00_399

import (
    "fmt"
)

//357. 计算各个位数不同的数字个数
func countNumbersWithUniqueDigits(n int) int {

    if n == 0 {
        return 1
    }

    res, muls := 10, 9

    for i := 1; i < n; i++ {
        muls *= 10 - i
        res += muls
    }

    return res
}

func main() {
    fmt.Println(countNumbersWithUniqueDigits(4))

}
