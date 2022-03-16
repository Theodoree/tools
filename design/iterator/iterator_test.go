package iterator

import (
    "fmt"
    "testing"
)

/* 迭代器
tip:
    常见场景: 复杂的迭代条件,一般使用循环就可以了,复杂场景需要这玩意,像go的hmap.
*/

func TestCreateUserIterator(t *testing.T) {
    iterator := CreateUserIterator([]User{{Id: 55}, {Id: 18}})

    for iterator.HasNext() {
        val,ok:=iterator.Next().(User)
        if !ok {
            t.Fatal(fmt.Sprintf("类型错误"))
        }
        println(val.Id)
    }
}
