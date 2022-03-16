package responsibility

import (
    "fmt"
    "testing"
)

/*
tip:
    应用场景:
        过滤敏感词,一个请求不确定需要使用哪一个函数处理
    职责链模式:多个处理器依次处理同一个请求。一个请求先经过A处理器处理，然后 再把请求传递给B处理器，B处理器处理完后再传递给C处理器，以此类推，形成一个链条。

*/



func TestNewFilterChain(t *testing.T) {
    chain:=NewFilterChain()


    chain.AddSubFilter(NewAdWordFilter())
    chain.AddSubFilter(NewsexWordFilter())

    postText:=`大扎好,我是渣渣辉,玩游戏无所谓,是兄弟就来砍我`

    fmt.Printf("过滤后:%s \n",chain.DoFilter(postText))



}

