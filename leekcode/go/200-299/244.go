package _00_299

import (
    "math"
)

/*
244. 最短单词距离 II

请设计一个类，使该类的构造函数能够接收一个单词列表。然后再实现一个方法，该方法能够分别接收两个单词 word1 和 word2，并返回列表中这两个单词之间的最短距离。您的方法将被以不同的参数调用 多次。

示例:
假设 words = ["practice", "makes", "perfect", "coding", "makes"]

输入: word1 = “coding”, word2 = “practice”
输出: 3
输入: word1 = "makes", word2 = "coding"
输出: 1
注意:
你可以假设 word1 不等于 word2, 并且 word1 和 word2 都在列表里。
*/

type WordDistance struct {
    arr []string
    m   map[string][]int
}

func Constructor(words []string) WordDistance {
    word := WordDistance{m: map[string][]int{}, arr: words}
    for k, v := range words {
        word.m[v] = append(word.m[v], k)
    }
    return word

}

func (this *WordDistance) Shortest(word1 string, word2 string) int {

    var min int
    var val int
    min = 0x666666
    for _, k := range this.m[word1] {

        for _, k1 := range this.m[word2] {
            val = int(math.Abs(float64(k - k1)))
            if val < min {
                min = val
            }
        }
    }

    return min
}
