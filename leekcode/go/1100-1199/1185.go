package _100_1199


/*
1185. 一周中的第几天

给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。

输入为三个整数：day、month 和 year，分别表示日、月、年。

您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。



示例 1：

输入：day = 31, month = 8, year = 2019
输出："Saturday"
示例 2：

输入：day = 18, month = 7, year = 1999
输出："Sunday"
示例 3：

输入：day = 15, month = 8, year = 1993
输出："Sunday"
*/

func dayOfTheWeek(day int, month int, year int) string {
    var (
        m      = strconv.Itoa(month)
        d      = strconv.Itoa(day)
        result string
    )

    if day < 10 {
        d = "0" + d
    }
    if month < 10 {
        m = "0" + m
    }

    // 这里使用内置的time包,如果不使用time包 求从1970从现在有多少天 然后%7
    t, _ := time.Parse(`2006-01-02`, fmt.Sprintf("%d-%s-%s", year, m, d))

    switch t.Weekday() {
    case time.Sunday:
        result = "Sunday"
    case time.Monday:
        result = "Monday"
    case time.Tuesday:
        result = "Tuesday"
    case time.Wednesday:
        result = "Wednesday"
    case time.Thursday:
        result = "Thursday"
    case time.Friday:
        result = "Friday"
    case time.Saturday:
        result = "Saturday"
    }

    return result

}

