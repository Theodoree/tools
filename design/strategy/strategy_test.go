package strategy

import "testing"

/*
tip:      策略模式:它解耦的是策略的定义、创建、使用这三部分
    使用场景:
        登录可以分为邮箱、账号、扫码登录,那么可以将登录抽象成这个策略,根据入参的类型选择一种处理策略,创建并且调用策略
*/

func TestGetPayFactory(t *testing.T) {
    InitDefaultFactory()
    payStrategy, err := GetPayStrategy(AliPayEnum)
    if err != nil {
        t.Fatal(err)
    }
    payStrategy.Pay(100)
}
