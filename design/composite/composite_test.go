package composite

import (
    "fmt"
    "testing"
)

/*
tip：
    场景:
        业务场景可以形成树状结构,组合模式将一组对象组织形成树状结构,以统一的接口抽象进行递归调用。

*/

func TestNewManager(t *testing.T) {

    li := NewManager("李经理")

    sun := NewSeller("孙xx", 7451)
    xu := NewSeller("徐xx", 11234)

    li.AddSeller(sun)
    li.AddSeller(xu)


    fmt.Printf("%s的总销售额: %d  \n",li.Name(),li.SellerStats())


}
