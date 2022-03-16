package _100_2199


/*
2144. 打折购买糖果的最小开销
一家商店正在打折销售糖果。每购买 两个 糖果，商店会 免费 送一个糖果。
免费送的糖果唯一的限制是：它的价格需要小于等于购买的两个糖果价格的 较小值 。
比方说，总共有 4 个糖果，价格分别为 1 ，2 ，3 和 4 ，一位顾客买了价格为 2 和 3 的糖果，那么他可以免费获得价格为 1 的糖果，但不能获得价格为 4 的糖果。
给你一个下标从 0 开始的整数数组 cost ，其中 cost[i] 表示第 i 个糖果的价格，请你返回获得 所有 糖果的 最小 总开销。
*/

func minimumCost(cost []int) int {
	sort.Ints(cost)

	var costSum int
	var cnt int
	min:=math.MaxInt64
	for i:=len(cost)-1;i>=0;i--{
		if cost[i] == -1 {
			continue
		}
		if cnt == 2 {
			cnt = 0
			continue
		}

		if cost[i] < min {
			min = cost[i]
		}
		costSum+=cost[i]
		cnt++
	}

	return costSum

}

