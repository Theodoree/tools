package adapter

import "testing"

/*
   tip:
      使用场景:用于补救设计上的缺陷
      出现接口不兼容场景:
      封装有缺陷的接口设计
      统一多个类的接口设计
      替换依赖的外部系统
      兼容老版本接口
      适配不同格式的数据
      适配器模式:适配器模式是一种事后的补救策略。适配器提供跟原始类不同的接口，而代理模式、装饰器模式提供的都是跟原始类相同的接口。
*/

func TestNewBankAccountAdapte(t *testing.T) {

    var wallet Wallet

    bankAccount := &bankAccount{balance: 884.53}

    wallet = NewBankAccountAdapte(bankAccount)

    println(wallet.Balance())

}
