package _00_399



/*
318. 最大单词长度乘积




题目描述
评论 (38)
题解(7)New
提交记录
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

示例 1:

输入: ["abcw","baz","foo","bar","xtfn","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "xtfn"。
示例 2:

输入: ["a","ab","abc","d","cd","bcd","abcd"]
输出: 4
解释: 这两个单词为 "ab", "cd"。
示例 3:

输入: ["a","aa","aaa","aaaa"]
输出: 0
解释: 不存在这样的两个单词。
*/

func maxProduct(words []string) int {

    var k, max int
    k = 0
    var arr = make([]int, 26)
    var flag bool
    for k < len(words) {
        arr = make([]int, 26)
        flag = false
        for _, v := range words[k] {
            arr[v-'a'] = 1
        }

        for j := k + 1; j < len(words); j++ {
            flag = false
            for _, v := range words[j] {
                if arr[v-'a'] == 1 {
                    flag = true
                    break
                }
            }

            if !flag {
                sum := len(words[j]) * len(words[k])
                if sum > max {
                    max = sum
                }
            }

        }

        k++
    }

    return max

}

// 位运算
func maxProduct(words []string) int {
    n:=len(words)
    arr:=make([]int,n)
    for i:=0;i<n;i++{
        len:=len(words[i])
        for j:=0;j<len;j++{
            arr[i] |=1<<(words[i][j]-'a')
        }
    }
    ans:=0
    for i:=0;i<n;i++{
        for j:=i+1;j<n;j++{
            if arr[i]&arr[j]==0{
                k:=len(words[i])*len(words[j])
                if ans<k{
                    ans=k
                }else{k=ans}
            }
        }
    }
    return ans
}
