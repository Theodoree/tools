package builder

import (
    "fmt"
    "testing"
)

/* 建造者模式
tip: 大致就是使用builder去做断言,减少断言对结构的侵入度


tip: 建造者使用场景:
    1.构造方法必填属性很多，需要检验
    2.类属性之间有依赖关系或者约束条件...
*/

func TestBuilder(t *testing.T) {
    user, _ := Builder().SetAge(15).SetNickName("alice").build()
    fmt.Printf("%#v \n", user)
}
