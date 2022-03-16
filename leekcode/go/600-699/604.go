package _00_699

/*
604. 迭代压缩字符串

对于一个压缩字符串，设计一个数据结构，它支持如下两种操作： next 和 hasNext。

给定的压缩字符串格式为：每个字母后面紧跟一个正整数，这个整数表示该字母在解压后的字符串里连续出现的次数。

next() - 如果压缩字符串仍然有字母未被解压，则返回下一个字母，否则返回一个空格。
hasNext() - 判断是否还有字母仍然没被解压。

注意：

请记得将你的类在 StringIterator 中 初始化 ，因为静态变量或类变量在多组测试数据中不会被自动清空。更多细节请访问 这里 。

示例：

StringIterator iterator = new StringIterator("L1e2t1C1o1d1e1");

iterator.next(); // 返回 'L'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 'e'
iterator.next(); // 返回 't'
iterator.next(); // 返回 'C'
iterator.next(); // 返回 'o'
iterator.next(); // 返回 'd'
iterator.hasNext(); // 返回 true
iterator.next(); // 返回 'e'
iterator.hasNext(); // 返回 false
iterator.next(); // 返回 ' '
*/

type StringIterator struct {
    s        string
    curIdx   int
    NumIndex int
    Less     int
}

func Constructor(compressedString string) StringIterator {
    var less int
    var i int
    for i = 1; i < len(compressedString); i++ {
        val := compressedString[i] - '0'
        if val >= 0 && val <= 9 {
            less = less * 10 + int(val)
        } else {
            break
        }

    }

    return StringIterator{
        s:        compressedString,
        curIdx:   0,
        Less:     less,
        NumIndex: i,
    }
}

func (this *StringIterator) Next() byte {
    if !this.HasNext() {
        return ' '
    }
    val := this.s[this.curIdx]
    this.Less--

    if this.Less == 0 {
        this.curIdx = this.NumIndex
        var less int
        var i int
        for i = this.curIdx + 1; i < len(this.s); i++ {
            val := this.s[i] - '0'
            if val >= 0 && val <= 9 {
                less =  less * 10 + int(val)
            } else {
                break
            }
        }
        this.NumIndex = i
        this.Less = less
    }

    return val
}

func (this *StringIterator) HasNext() bool {
    return !(this.curIdx >= len(this.s))
}
