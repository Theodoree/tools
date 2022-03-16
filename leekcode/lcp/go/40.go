package _go

import "sort"

/*
LCP 40. 心算挑战
「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。若不存在获取有效得分的卡牌方案，则返回 0。
*/
func maxmiumScore(cards []int, cnt int) int {
	sort.Ints(cards)

	var sum int
	var n [2]int
	n[0] = 0
	n[1] = 0
	for i:=len(cards)-1;i>=len(cards)-cnt;i--{
		if cards[i]%2 == 0 && ( n[0] == 0 || cards[i] < n[0]){
			n[0] = cards[i]
		}else if cards[i]%2 == 1 && ( n[1] == 0 ||cards[i] < n[1]){
			n[1] = cards[i]
		}
		sum+=cards[i]
	}

	if sum%2 == 0 {
		return sum
	}

	max:= func(i,j int) int {
		if i > j {
			return i
		}
		return j
	}

	var maxArr [2]int
	for i:=0;i<len(cards)-cnt;i++{
		if cards[i]%2 == 0 && cards[i] > maxArr[0]{
			maxArr[0] = cards[i]
		}else if cards[i]%2 == 1 && cards[i] > maxArr[1]{
			maxArr[1] = cards[i]
		}
	}

	if n[0] > 0 && maxArr[1] > 0 && n[1] > 0 && maxArr[0] > 0{
		sum = max(sum-n[0]+maxArr[1],sum-n[1]+maxArr[0])
	}else if n[0] > 0 && maxArr[1] > 0 {
		sum = sum-n[0]+maxArr[1]
	}else if n[1] > 0 && maxArr[0] > 0 {
		sum = sum-n[1]+maxArr[0]
	}

	if sum %2 == 0 {
		return sum
	}

	return 0
}

