package _100_1199



/*
1189. “气球” 的最大数量

给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。

字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。



示例 1：



输入：text = "nlaebolko"
输出：1
示例 2：



输入：text = "loonbalxballpoon"
输出：2
示例 3：

输入：text = "leetcode"
输出：0
*/
func maxNumberOfBalloons(text string) int {
    str1 := `balloon`
    var arr [26]int
    var arr1 [26]int

    for _, v := range str1 {
        arr[v-'a']++
    }
    var cnt int

    for _, v := range text {
        arr1[v-'a']++
    }

lp:
    for {

        for i := 0; i < 26; i++ {
            arr1[i] -= arr[i]
            if arr1[i] < 0 {
                break lp
            }

        }
        cnt++

    }

    return cnt

}

