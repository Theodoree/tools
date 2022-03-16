package _00_299


/*
251. 展开二维向量

请设计并实现一个能够展开二维向量的迭代器。该迭代器需要支持 next 和 hasNext 两种操作。、

示例：

Vector2D iterator = new Vector2D([[1,2],[3],[4]]);

iterator.next(); // 返回 1
iterator.next(); // 返回 2
iterator.next(); // 返回 3
iterator.hasNext(); // 返回 true
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 4
iterator.hasNext(); // 返回 false


注意：

请记得 重置 在 Vector2D 中声明的类变量（静态变量），因为类变量会 在多个测试用例中保持不变，影响判题准确。请 查阅 这里。
你可以假定 next() 的调用总是合法的，即当 next() 被调用时，二维向量总是存在至少一个后续元素。

*/

type Vector2D struct {
    val       [][]int
    curIdx    int
    curIdxCol int
}

func Constructor(v [][]int) Vector2D {
    return Vector2D{
        val:       v,
        curIdx:    0,
        curIdxCol: 0,
    }

}

func (this *Vector2D) Next() int {
    for this.curIdx < len(this.val) && len(this.val[this.curIdx]) == 0 {
        this.curIdx++
    }
    if !this.HasNext() {
        return 0
    }
    val := this.val[this.curIdx][this.curIdxCol]

    if this.curIdxCol+1 > len(this.val[this.curIdx])-1 {
        this.curIdxCol = 0
        this.curIdx++
    } else {
        this.curIdxCol++
    }

    return val

}

func (this *Vector2D) HasNext() bool {
    for this.curIdx < len(this.val) && len(this.val[this.curIdx]) == 0 {
        this.curIdx++
    }
    if this.curIdx >= len(this.val) {
        return false
    }

    return true
}

