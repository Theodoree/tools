package facade

import "testing"

/*
tip: 门面模式
       :
        解决子接口粒度过细的问题,一个业务(分布式业务)可能需要多个接口配合调用,当如果从客户端调用,多次访问会导致网络开支过大。
    场景:
    1.解决易用性
        linux提供的系统调用机制(最少知识原则、门面模式)
    2.解决性能问题
        多次调用导致的网络连接开支


*/


func TestNewController(t *testing.T) {

    controller:=NewController()
    controller.AddSku(123,777)
}

