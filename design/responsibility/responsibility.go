package responsibility

import "strings"

// Chain Of Responsibility Design Pattern

// 比如想按照Gof的定义,那么在接口中定义一个布尔值表示是否能够处理即可。
// handler( val ) (val,executed)
// 如果executed等于false那么继续传递,使用下一个执行函数
type FilterHandler interface {
    filter(postContext string) string
}

type FilterChain interface {
    DoFilter(postContext string)  string
    AddSubFilter(filter FilterHandler)
}

func NewFilterChain() FilterChain {
    return &filterChain{}
}

type filterChain struct {
    arr []FilterHandler
}

func (chain *filterChain) DoFilter(postContext string)  string {

    for _, v := range chain.arr {
        postContext =  v.filter(postContext)
    }

    return postContext
}

func (chain *filterChain) AddSubFilter(filter FilterHandler) {
    // 当然这里可能会有问题,demo嘛,习惯就好,这个位置如果增加一个map过滤会更好。
    chain.arr = append(chain.arr, filter)
}

// 因为下面的结构都没有使用内部变量,所以就不使用指针接收者,因为没必要,这个struct只是用来区分空间
type adWordFilter struct{}
func (filter adWordFilter) filter(postContext string)  string {
    // do something
    postContext = strings.Replace(postContext, "贪玩蓝月", "", -1)
    postContext = strings.Replace(postContext, "渣渣辉", "", -1)
    return postContext
}
func NewAdWordFilter() FilterHandler {
    return adWordFilter{}
}


type sexWordFilter struct{}
func (f sexWordFilter) filter(postContext string) string {
    // 其实想这里可以考虑用那种分词库去过滤,大概职责链就是这么个东西。
    postContext = strings.Replace(postContext, "曹", "", -1)
    postContext = strings.Replace(postContext, "艹", "", -1)
    postContext = strings.Replace(postContext, "操", "", -1)
    return postContext
}
func NewsexWordFilter() FilterHandler {
    return sexWordFilter{}
}
