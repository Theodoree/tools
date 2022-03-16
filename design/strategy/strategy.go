package strategy

import "errors"

// 1. 策略的定义
type PayStrategy interface {
    Pay(amount uint64)

    init(config map[string]string) PayStrategy
}

type Type uint

const (
    AliPayEnum Type = iota + 1
    WeChatEnum
    BankEnum
)

// 2. 策略的创建
type payFactory struct {
    strategy map[Type]PayStrategy
}

func (p *payFactory) CreateStrategy(enum Type) (PayStrategy, error) {
    if val, ok := p.strategy[enum]; ok {
        return val, nil
    }

    return nil, unKnownEnum
}

var unKnownEnum = errors.New("未知的策略类型")
var needInit = errors.New("需要初始化")

var defaultFactory *payFactory

func InitDefaultFactory() {
    // 初始化默认工厂 可以在这里设置
    defaultFactory = &payFactory{strategy: make(map[Type]PayStrategy)}
    defaultFactory.strategy[AliPayEnum] = Alipay{}.init(nil)
    defaultFactory.strategy[WeChatEnum] = WechatPay{}.init(nil)
    defaultFactory.strategy[BankEnum] = BankPay{}.init(nil)

}

func GetPayStrategy(enum Type) (PayStrategy, error) {
    if defaultFactory == nil {
        return nil, needInit
    }
    return defaultFactory.CreateStrategy(enum)
}

// 3.策略的实现
type Alipay struct {
}

func (p *Alipay) Pay(amount uint64) {
    println("alipay")
}

func (p Alipay) init(config map[string]string) PayStrategy {
    return &Alipay{}
}

type WechatPay struct {
}

func (p *WechatPay) Pay(amount uint64) {
    println("WechatPay")

}

func (p WechatPay) init(config map[string]string) PayStrategy {
    return &WechatPay{}
}

type BankPay struct {
}

func (p *BankPay) Pay(amount uint64) {
    println("BankPay")

}

func (p BankPay) init(config map[string]string) PayStrategy {
    return &BankPay{}
}
