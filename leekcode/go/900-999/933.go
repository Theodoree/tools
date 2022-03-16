package _00_999



/*
933. 最近的请求次数

写一个 RecentCounter 类来计算最近的请求。

它只有一个方法：ping(int t)，其中 t 代表以毫秒为单位的某个时间。

返回从 3000 毫秒前到现在的 ping 数。

任何处于 [t - 3000, t] 时间范围之内的 ping 都将会被计算在内，包括当前（指 t 时刻）的 ping。

保证每次对 ping 的调用都使用比之前更大的 t 值。

示例：

输入：inputs = ["RecentCounter","ping","ping","ping","ping"], inputs = [[],[1],[100],[3001],[3002]]
输出：[null,1,2,3,3]

*/

type RecentCounter struct {
    f    []int
}

func Constructor() RecentCounter {
    return RecentCounter{make([]int, 0, 10000)}
}

func (this *RecentCounter) Ping(t int) int {
    this.f = append(this.f, t)

    for this.f[0] < t - 3000 {
        this.f = this.f[1:]
    }

    return len(this.f)
}