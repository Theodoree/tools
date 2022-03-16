package decorator

import (
    "testing"
)

/*
tip: 装饰器模式: 增强部分功能
    使用场景:如果设计者不需要用户关 注是否使用缓存功能，要隐藏实现细节，也就是说用户只能看到和使用代理类，那么就使 用proxy模式
    ;反之，如果设计者需要用户自己决定是否使用缓存的功能，需要用户自己新 建原始对象并动态添加缓存功能，那么就使用decorator模式。
    装饰器模式:装饰者模式在不改变原始类接口的情况下，对原始类功能进行增强，并且支持 多个装饰器的嵌套使用。
*/

func TestNewThreadPool(t *testing.T) {
    pool := NewThreadPool(10)

    for i:=0;i<10000;i++{
        pool.SendNewFunc(func() {
            println(1)
        })
    }

}
