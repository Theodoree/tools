package bridge

import "testing"

/* 桥接模式
   tip:将抽象和实现解耦，让它们可以独立变化
   tip:
    使用场景:
        一个功能可能会有多种实现类,比如说支付,钱包支付 移动支付 银行卡支付。
        数据库 mysql mongo tidb .....
   tip:
      与策略模式相比,桥接模式主要是体现在于接口隔离的原则,桥接仅提供基本操作
    桥接模式:桥接模式的目的是将接口部分和实现部分分离，从而让它们可以较为容易、也相 对独立地加以改变。
*/

func TestNewService(t *testing.T) {
    s, err := NewService("mongo:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("mysql:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("oracle:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()

    s, err = NewService("123:12312312")
    if err != nil {
        t.Fatal(err)
    }
    s.Open()
}
