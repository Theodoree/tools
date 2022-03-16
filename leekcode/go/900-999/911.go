package _00_999


/*
911. 在线选举
给你两个整数数组 persons 和 times 。在选举中，第 i 张票是在时刻为 times[i] 时投给候选人 persons[i] 的。

对于发生在时刻 t 的每个查询，需要找出在 t 时刻在选举中领先的候选人的编号。

在 t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

实现 TopVotedCandidate 类：

TopVotedCandidate(int[] persons, int[] times) 使用 persons 和 times 数组初始化对象。
int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。

*/
type TopVotedCandidate struct {
	times   []int
	maxPer  []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	mm := make([]int, 0, len(persons))
	maxPerson := 0
	m := make(map[int]int, len(persons))
	for _, last := range persons {
		m[last]++
		if m[last] >= m[maxPerson] {
			maxPerson = last
		}
		mm = append(mm, maxPerson)
	}

	return TopVotedCandidate{
		times:   times,
		maxPer:  mm,
	}
}

func (this *TopVotedCandidate) Q(t int) int {
	i := this.getTopIndex(t)
	return this.maxPer[i]
}

func (this *TopVotedCandidate) getTopIndex(t int) int {
	l := 0
	r := len(this.times) - 1
	for l <= r {
		m := (l + r) / 2
		if this.times[m] == t {
			return m
		} else if this.times[m] < t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l - 1
}
