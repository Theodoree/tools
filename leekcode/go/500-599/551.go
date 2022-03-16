package _00_599


/*
551. 学生出勤记录 I

给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：

'A' : Absent，缺勤
'L' : Late，迟到
'P' : Present，到场
如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。

你需要根据这个学生的出勤记录判断他是否会被奖赏。

示例 1:

输入: "PPALLP"
输出: True
示例 2:

输入: "PPALLL"
输出: False
*/

func checkRecord(s string) bool {

    var A bool
    var isLate int

    switch len(s) > 0 {
    case s[0] == 'A':
        A = true
    case s[0] == 'L':
        isLate = 1
    }

    for i := 1; i < len(s); i++ {

        if s[i] == 'L' && isLate == 2 {
            return false
        } else if s[i] == 'L' {
            isLate++
        } else {
            isLate = 0
        }

        if s[i] == 'A' && !A {
            A = true
        } else if s[i] == 'A' {
            return false
        }

    }

    return true

}