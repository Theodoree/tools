package _00_599


/*
599. 两个列表的最小索引总和

假设Andy和Doris想在晚餐时选择一家餐厅，并且他们都有一个表示最喜爱餐厅的列表，每个餐厅的名字用字符串表示。

你需要帮助他们用最少的索引和找出他们共同喜爱的餐厅。 如果答案不止一个，则输出所有答案并且不考虑顺序。 你可以假设总是存在一个答案。

示例 1:

输入:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
输出: ["Shogun"]
解释: 他们唯一共同喜爱的餐厅是“Shogun”。
示例 2:

输入:
["Shogun", "Tapioca Express", "Burger King", "KFC"]
["KFC", "Shogun", "Burger King"]
输出: ["Shogun"]
解释: 他们共同喜爱且具有最小索引和的餐厅是“Shogun”，它有最小的索引和1(0+1)。
*/

func findRestaurant(list1 []string, list2 []string) []string {
    m := make(map[string]int)

    for k, v := range list1 {
        m[v] = k + (len(list1) + len(list2))
    }

    var min, result int
    min = 1e9
    for k, v := range list2 {
        if _, ok := m[v]; ok {
            result = m[v] + k - (len(list1) + len(list2))
            m[v] = result
            if result < min {
                min = result
            }
        }
    }

    var resultStr []string
    for k, v := range m {
        if v == min {
            resultStr = append(resultStr, k)
        }
    }
    return resultStr
}


func findRestaurant(list1 []string, list2 []string) []string {
    nameIndexMap := make(map[string]int)
    for i, name := range list1 {
        nameIndexMap[name] = i
    }
    var res []string
    currentMin := 99999
    for i, name := range list2 {
        _, ok := nameIndexMap[name]
        if ok {
            if nameIndexMap[name] + i < currentMin {
                currentMin = nameIndexMap[name] + i
                res = []string{name}
            } else if nameIndexMap[name] + i == currentMin {
                res = append(res, name)
            }
        }
    }
    return res
}