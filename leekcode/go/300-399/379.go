package _00_399

/*
379. 电话目录管理系统

设计一个电话目录管理系统，让它支持以下功能：

get: 分配给用户一个未被使用的电话号码，获取失败请返回 -1
check: 检查指定的电话号码是否被使用
release: 释放掉一个电话号码，使其能够重新被分配
示例:

// 初始化电话目录，包括 3 个电话号码：0，1 和 2。
PhoneDirectory directory = new PhoneDirectory(3);

// 可以返回任意未分配的号码，这里我们假设它返回 0。
directory.get();

// 假设，函数返回 1。
directory.get();

// 号码 2 未分配，所以返回为 true。
directory.check(2);

// 返回 2，分配后，只剩一个号码未被分配。
directory.get();

// 此时，号码 2 已经被分配，所以返回 false。
directory.check(2);

// 释放号码 2，将该号码变回未分配状态。
directory.release(2);

// 号码 2 现在是未分配状态，所以返回 true。
directory.check(2);
*/
type PhoneDirectory struct {
    n, i, ir int
    rec      []int
    state    []int
}

func Constructor(maxNumbers int) PhoneDirectory {
    return PhoneDirectory{maxNumbers, 0, -1, make([]int, maxNumbers), make([]int, maxNumbers)}
}

func (this *PhoneDirectory) Get() int {
    out := -1
    if this.i < this.n {
        out = this.i
        this.i++
    } else if this.ir >= 0 {
        out = this.rec[this.ir]
        this.ir--
    }
    if out != -1 {
        this.state[out] = 1
    }
    return out
}

func (this *PhoneDirectory) Check(number int) bool {
    return this.state[number] == 0
}

func (this *PhoneDirectory) Release(number int) {
    if this.state[number] == 0 {
        return
    }
    this.ir++
    this.rec[this.ir] = number
    this.state[number] = 0
}

/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */
