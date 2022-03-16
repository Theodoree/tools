package observer

import "testing"

/*
tip:
    desc:
        类似发布-订阅模型,一对多的关系,一条信息可以被n个订阅消费。可以同步也可以异步。
        大致思路就是通过接口抽象,将某一部分嵌入的业务,作为一个接口抽象数组载入业务中,从而解耦合,如果新增处理函数,只需要实现该接口,并且增加到数组中即可。
        :
        具体到观察者模式，它将观察者和被观察者解耦。借助设计模式,将一大坨代码拆分为职责更细的小类,以此满足开闭、高内聚低耦合的特性。从而降低代码的复杂度和可拓展性。
    使用场景:
        邮件订阅、消息队列、电商系统(订单达成,通知下游子系统)
*/

func TestNewController(t *testing.T) {
    controller:=NewController()
    controller.SetRegObserver([]RegObserver{&RegNotificationObserver{}})

    userId:=controller.Register("13479180942","110")
    println(userId)
}