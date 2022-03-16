package _00_599

import "math"

/*
504. 七进制数
给定一个整数，将其转化为7进制，并以字符串形式输出。

示例 1:

输入: 100
输出: "202"
示例 2:

输入: -7
输出: "-10"
*/

func convertToBase7(num int) string {

    if num == 0 {
        return "0"
    }

    isNegative := num < 0
    num = int(math.Abs(float64(num)))

    var result string
    var r int
    for num > 0 {
        r = num % 7
        num /= 7
        result = strconv.Itoa(r) + result

    }

    //100
    //202
    if isNegative {
        result = "-" + result
    }

    return result
}
