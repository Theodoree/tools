package _000_1099

import (
    "container/heap"
    "sort"
)

/*
1086. 前五科的均分


给你一个不同学生的分数列表，请按 学生的 id 顺序 返回每个学生 最高的五科 成绩的 平均分。

对于每条 items[i] 记录， items[i][0] 为学生的 id，items[i][1] 为学生的分数。平均分请采用整数除法计算。



示例：

输入：[[1,91],[1,92],[2,93],[2,97],[1,60],[2,77],[1,65],[1,87],[1,100],[2,100],[2,76]]
输出：[[1,87],[2,88]]
解释：
id = 1 的学生平均分为 87。
id = 2 的学生平均分为 88.6。但由于整数除法的缘故，平均分会被转换为 88。


提示：

1 <= items.length <= 1000
items[i].length == 2
学生的 ID 在 1 到 1000 之间
学生的分数在 1 到 100 之间
每个学生至少有五个分数

*/



func highFive(items [][]int) [][]int {
    val := make(map[int]*IntHeap)
    for i := 0; i < len(items); i++ {
        if _, ok := val[items[i][0]]; !ok {
            h := &IntHeap{}
            heap.Init(h)
            val[items[i][0]] = h
        }
        heap.Push(val[items[i][0]], items[i][1])
        if val[items[i][0]].Len() > 5 {
            heap.Pop(val[items[i][0]])
        }
    }
    ret := make([][]int, 0)
    for k, h := range val {
        score := 0
        for h.Len() > 0 {
            score += heap.Pop(h).(int)
        }
        ret = append(ret, []int{k, score / 5})
    }
    sort.Slice(ret, func(i, j int) bool {
        return ret[i][0] < ret[j][0]
    })
    return ret
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}