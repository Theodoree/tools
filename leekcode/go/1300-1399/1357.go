package _300_1399


/*
1357. 每隔 n 个顾客打折
超市里正在举行打折活动，每隔 n 个顾客会得到 discount 的折扣。

超市里有一些商品，第 i 种商品为 products[i] 且每件单品的价格为 prices[i] 。

结账系统会统计顾客的数目，每隔 n 个顾客结账时，该顾客的账单都会打折，折扣为 discount （也就是如果原本账单为 x ，那么实际金额会变成 x - (discount * x) / 100 ），然后系统会重新开始计数。

顾客会购买一些商品， product[i] 是顾客购买的第 i 种商品， amount[i] 是对应的购买该种商品的数目。

请你实现 Cashier 类：

Cashier(int n, int discount, int[] products, int[] prices) 初始化实例对象，参数分别为打折频率 n ，折扣大小 discount ，超市里的商品列表 products 和它们的价格 prices 。
double getBill(int[] product, int[] amount) 返回账单的实际金额（如果有打折，请返回打折后的结果）。返回结果与标准答案误差在 10^-5 以内都视为正确结果。

*/

type Cashier struct {
    cnt      int
    m        map[int]int
    discount int // 打折率
    n        int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
    c := Cashier{
        cnt:      0,
        m:        make(map[int]int),
        discount: discount,
        n:        n,
    }

    for i := 0; i < len(products); i++ {
        c.m[products[i]] = prices[i]
    }

    return c

}

func (c *Cashier) GetBill(product []int, amount []int) float64 {
    var money int

    for i := 0; i < len(product); i++ {

        money += c.m[product[i]] * amount[i]
    }

    c.cnt++
    if c.cnt%c.n == 0 {
        return float64(money) * (1.0 - float64(c.discount)/100)

    }
    return float64(money)
}
