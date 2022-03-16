package proxy

import "time"

// Proxy Design Pattern
// 洋葱模型

type context struct {
    index int
    chain handlerChain
}

func (c *context) Next() {
    if c.index > len(c.chain) {
        return
    }
    oldIndex := c.index
    c.index++
    c.chain[oldIndex](c)
}

type handler = func(*context)

type handlerChain []handler


func recordExecutionTime(context *context) {
    n := time.Now()
    context.Next()
    println(time.Now().Sub(n))

}

func PrintEnv1(context *context) {
    println("Env1")
}

