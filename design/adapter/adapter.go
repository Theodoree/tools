package adapter

type Wallet interface {
    Balance() uint64 // 分
}

type BankAccount interface {
    Balance() float64
}

type bankAccount struct {
    balance float64
}

func (b *bankAccount) Balance() float64 {
    return b.balance
}

// 适配器
type BankAccountAdapte struct {
    *bankAccount
}

func (b *BankAccountAdapte) Balance() uint64 {
    return uint64(b.balance * 100)
}

func NewBankAccountAdapte(bankAccount *bankAccount) *BankAccountAdapte {
    return &BankAccountAdapte{bankAccount}
}
