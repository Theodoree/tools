package _100_1199

/*
1176. 健身计划评估
你的好友是一位健身爱好者。前段日子，他给自己制定了一份健身计划。现在想请你帮他评估一下这份计划是否合理。
他会有一份计划消耗的卡路里表，其中 calories[i] 给出了你的这位好友在第 i 天需要消耗的卡路里总量。
为了更好地评估这份计划，对于卡路里表中的每一天，你都需要计算他 「这一天以及之后的连续几天」 （共 k 天）内消耗的总卡路里 T：
如果 T < lower，那么这份计划相对糟糕，并失去 1 分；
如果 T > upper，那么这份计划相对优秀，并获得 1 分；
否则，这份计划普普通通，分值不做变动。
请返回统计完所有 calories.length 天后得到的总分作为评估结果。
注意：总分可能是负数。
*/
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	presum := make([]int, len(calories)+1)
	presum[0] = 0
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + calories[i-1]
	}
	var point int
	for i := k; i < len(presum); i++ {
		t := presum[i]
		if i > k {
			t = presum[i] - presum[i-k]
		}
		if t < lower {
			point--
		} else if t > upper {
			point++
		}
	}

	return point

}
