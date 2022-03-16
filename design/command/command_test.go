package command

import (
    "testing"
    "time"
)


/* 命令模式
tip:
    将处理过程抽象为一个command,统一的向外接受流量,根据类型的不同使用不同的command,与策略模式不同的一点在于,策略模式最终是要达成一个目的,比如说支付会有支付宝、银行卡、微信支付。
desc:
*/


func TestNewClient(t *testing.T) {
    srv:=NewService()

    srv.Start()


    client:=srv.Connection("192.168.0.1")

    client.PushEvents(&event{
        Action: printCmd,
        Data:   "哈哈哈哈哈哈哈",
    })

    client.PushEvents(&event{
        Action: printCmd,
        Data:   "11111111",
    })
    client.PushEvents(&event{
        Action: printCmd,
        Data:   "bbbbbbb",
    })

    time.Sleep(time.Second*3)
}
