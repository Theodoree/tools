package _100_1199




/*
1169. 查询无效交易

如果出现下述两种情况，交易 可能无效：

交易金额超过 ¥1000
或者，它和另一个城市中同名的另一笔交易相隔不超过 60 分钟（包含 60 分钟整）
每个交易字符串 transactions[i] 由一些用逗号分隔的值组成，这些值分别表示交易的名称，时间（以分钟计），金额以及城市。

给你一份交易清单 transactions，返回可能无效的交易列表。你可以按任何顺序返回答案。



示例 1：

输入：transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
输出：["alice,20,800,mtv","alice,50,100,beijing"]
解释：第一笔交易是无效的，因为第二笔交易和它间隔不超过 60 分钟、名称相同且发生在不同的城市。同样，第二笔交易也是无效的。
示例 2：

输入：transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
输出：["alice,50,1200,mtv"]
示例 3：

输入：transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
输出：["bob,50,1200,mtv"]
*/
type transaction struct {
    Nick      string
    TimeStamp int
    Money     int
    City      string
    Str       string
}

func invalidTransactions(transactions []string) []string {

    var t = make(map[string][]*transaction)

    for _, v := range transactions {
        val := strings.Split(v, ",")
        if len(val) != 4 {
            continue
        }
        timestamp, _ := strconv.Atoi(val[1])
        money, _ := strconv.Atoi(val[2])
        order := &transaction{
            Nick:      val[0],
            TimeStamp: timestamp,
            Money:     money,
            City:      val[3],
            Str:       v,
        }

        t[order.Nick] = append(t[order.Nick], order)
    }
    /*
       交易金额超过 ¥1000
       或者，它和另一个城市中同名的另一笔交易相隔不超过 60 分钟（包含 60 分钟整）
    */

    var result []string
    var resultMap = make(map[string]struct{})
    for _, v1 := range t {

        for _, trans := range v1 {
            if trans.Money >= 1000 {
                resultMap[trans.Str] = struct{}{}
            }

            for _, v := range v1 {
                if abs(v.TimeStamp-trans.TimeStamp) <= 60 && v.City != trans.City {
                    resultMap[v.Str] = struct{}{}
                    resultMap[trans.Str] = struct{}{}
                }
            }

        }
    }

    for k, _ := range resultMap {
        result = append(result, k)
    }

    return result
}

func abs(num int) int {
    if num < 0 {
        return - num
    }
    return num
}
