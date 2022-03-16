package _go

type RecentCounter struct {
	f    []int
}

func Constructor() RecentCounter {
	return RecentCounter{make([]int, 0, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.f = append(this.f, t)

	for this.f[0] < t - 3000 {
		this.f = this.f[1:]
	}

	return len(this.f)
}