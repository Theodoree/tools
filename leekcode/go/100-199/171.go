package _00_199


/*
171. Excel表列序号

给定一个Excel表格中的列名称，返回其相应的列序号。


*/
func titleToNumber(s string) int {

    var sum, result int
    var cnt int
    for i := len(s) - 1; i >= 0; i-- {
        result = int(s[i] - 'A' + 1)
        if sum == 0 {
            sum = result
        } else {
            for i := 0; i < cnt; i++ {
                result *= 26
            }
            sum += result

        }
        cnt++
    }

    return sum
}