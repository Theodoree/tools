package _00_599


/*
557. 反转字符串中的单词 III

给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:
输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
*/

func reverseWords(s string) string {
    var result []string
    var current string

    for _, v := range strings.Split(s, ` `) {
        current = ""
        for j := len(v) - 1; j >= 0; j-- {
            current += string(v[j])
        }
        result = append(result, current)
    }

    return  strings.Join(result,` `)
}

func reverseWords(s string) string {
    n := len(s)
    if n <= 1 {
        return ""
    }
    ss := []byte(s)
    first, last := -1, -1
    for i, c := range ss {
        if c != ' ' {
            if i == 0 || ss[i-1] == ' ' {
                first = i
            }
            if i == n-1 || ss[i+1] == ' ' {
                last = i

                for ;first < last; first, last = first+1, last-1 {
                    ss[first], ss[last] = ss[last], ss[first]
                }
            }
        }
    }

    return string(ss)
}