package proxy

import "testing"

/*

 tip:
  使用场景:
     洋葱模型,一般使用在web中间件中,鉴权、限流、日志
  代理模式:
        代理模式在不改变原始类接口的条件下，为原始类定义一个代理类，主要目的是控制访问，而非加强功能，这是它跟装饰器模式最大的不同。
*/
func TestProxy(t *testing.T) {

    var chain handlerChain
    chain = append(chain,recordExecutionTime)
    chain = append(chain,PrintEnv1)
    var c context
    c.chain =chain
    c.Next()


}
