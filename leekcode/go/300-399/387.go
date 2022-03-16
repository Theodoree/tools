package _00_399

//387. 字符串中的第一个唯一字符
func firstUniqChar1(s string) int {

    var filerMap = make(map[int32]int)
    f := make([]bool, len(s), len(s))

    for i, key := range s {
        if index, ok := filerMap[key]; ok {
            f[i] = true
            f[index] = true
        } else {
            filerMap[key] = i
        }
    }

    for k, v := range f {
        if v == false {
            return k
        }
    }
    return -1
}



func firstUniqChar(source string) int {
    counts := make([]int, 26)

    for _, item := range source {
        counts[item -'a']++
    }
    for index, item := range source {
        if counts[item - 'a'] == 1 {
            return index
        }
    }
    return -1
}

