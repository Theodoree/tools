package _000_1099

import "sort"

/*
1029. 两地调度

公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。

返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。



示例：

输入：[[10,20],[30,200],[400,50],[30,20]]
输出：110
解释：
第一个人去 A 市，费用为 10。
第二个人去 A 市，费用为 30。
第三个人去 B 市，费用为 50。
第四个人去 B 市，费用为 20。

最低总费用为 10 + 30 + 50 + 20 = 110，每个城市都有一半的人在面试。




cost a = TeamA -> A  + TeamB -> B
cost b = TeamA -> B  + TeamB -> A

求 min(a)

*/
func twoCitySchedCost(costs [][]int) int {

    sort.Slice(costs, func(i, j int) bool {
        return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
    })

    var result int
    var left, right int
    right = len(costs) - 1
    for left < right {
        result += costs[left][0]
        result += costs[right][1]

        left++
        right--
    }

    return result
}
