package _000_1099

import "fmt"

/*
1010. 总持续时间可被 60 整除的歌曲

在歌曲列表中，第 i 首歌曲的持续时间为 time[i] 秒。

返回其总持续时间（以秒为单位）可被 60 整除的歌曲对的数量。形式上，我们希望索引的数字  i < j 且有 (time[i] + time[j]) % 60 == 0。
*/
func numPairsDivisibleBy60(time []int) int {
    m := make([]int, 60)
    cnt := 0
    for _, n := range time {
        fmt.Println(m[0])
        if n%60 == 0 {
            cnt += m[0]
        } else {
            cnt += m[60-n%60]
        }
        m[n%60]++
    }
    return cnt
}