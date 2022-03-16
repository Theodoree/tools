package _200_1299

import "sort"

/*
1244. 力扣排行榜
新一轮的「力扣杯」编程大赛即将启动，为了动态显示参赛者的得分数据，需要设计一个排行榜 Leaderboard。
请你帮忙来设计这个 Leaderboard 类，使得它有如下 3 个函数：
addScore(playerId, score)：
假如参赛者已经在排行榜上，就给他的当前得分增加 score 点分值并更新排行。
假如该参赛者不在排行榜上，就把他添加到榜单上，并且将分数设置为 score。
top(K)：返回前 K 名参赛者的 得分总和。
reset(playerId)：将指定参赛者的成绩清零（换句话说，将其从排行榜中删除）。题目保证在调用此函数前，该参赛者已有成绩，并且在榜单上。
请注意，在初始状态下，排行榜是空的。
*/

type Leaderboard struct {
	rank []int
	m map[int]int
}


func Constructor() Leaderboard {
	return Leaderboard{m: make(map[int]int)}
}


func (l *Leaderboard) AddScore(playerId int, score int)  {
	if _,ok:=l.m[playerId];!ok{
		l.rank = append(l.rank,playerId)
		l.m[playerId]=score
	}else{
		l.m[playerId]+=score
	}
}


func (l *Leaderboard) Top(K int) int {
	sort.Slice(l.rank, func(i, j int) bool {
		return l.m[l.rank[i]] > l.m[l.rank[j]]
	})
	var result int
	for i:=0;i<K;i++{
		result +=l.m[l.rank[i]]
	}

	return result

}


func (l *Leaderboard) Reset(playerId int)  {
	l.m[playerId] = 0
}



