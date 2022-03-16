package _00_499

import "strconv"

/*
415. 字符串相加

给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/
func addStrings(num1 string, num2 string) string {

    var n1R = len(num1) - 1
    var n2R = len(num2) - 1
    var result string
    var r, isLargeThanTen int
    for {
        if n1R >= 0 && n2R >= 0 {
            r = int((num1[n1R]+num2[n2R])-'1'*2+2) + isLargeThanTen
            n1R--
            n2R--
        } else if n1R >= 0 {
            r = int(num1[n1R]-'1'+1) + isLargeThanTen
            n1R--
        } else if n2R >= 0 {
            r = int(num2[n2R]-'1'+1) + isLargeThanTen
            n2R--
        } else {
            break
        }


        if r >= 10 {
            isLargeThanTen = 1
            r %= 10
        } else {
            isLargeThanTen = 0
        }
        result = strconv.Itoa(r) + result


    }
    if isLargeThanTen > 0 {
        result = strconv.Itoa(isLargeThanTen) + result
    }


    return result

}
