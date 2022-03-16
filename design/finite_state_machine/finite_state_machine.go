package finite_state_machine

// 场景: 模拟人生

type SimsController interface {
	Eat()
	Drink()
	WatchTv()
	Grow()
}

type StatusEnum int

const (
	Child StatusEnum = iota + 1
	Teenager
	Younger
	Older
)

// 分支法

func NewBranchSims() SimsController {
	return &BranchSims{
		StatusEnum: Child,
	}
}

type BranchSims struct {
	StatusEnum StatusEnum
}

func (b *BranchSims) Eat() {
	switch b.StatusEnum {
	case Child:
		println("小孩子:喝稀饭")
	case Teenager:
		println("青少年:吃米线")
	case Younger:
		println("年轻人:吃营养餐")
	case Older:
		println("老年人:吃饭")
	}

}

func (b *BranchSims) Drink() {
	switch b.StatusEnum {
	case Child:
		println("小孩子:喝牛奶")
	case Teenager:
		println("青少年:喝提神补脑液")
	case Younger:
		println("年轻人:喝水")
	case Older:
		println("老年人:喝脑x金")
	}

}

func (b *BranchSims) WatchTv() {
	switch b.StatusEnum {
	case Child:
		println("小孩子:看卡通片")
	case Teenager:
		println("青少年:看偶像剧")
	case Younger:
		// tip:如果把b倒过来也行
		println("年轻人:看b站、看直播")
	case Older:
		println("老年人:看年代剧")
	}

}

func (b *BranchSims) Grow() {
	if b.StatusEnum < Older {
		b.StatusEnum++
	}
}

// 查表法

// 可以改成function,怎么都行,这里就先简单的写一下
var statusTable = [][]string{
	{}, // nil
	{"小孩子:喝稀饭", "小孩子:喝牛奶", "小孩子:看卡通片"},    // Child
	{"青少年:吃米线", "青少年:喝提神补脑液", "青少年:看偶像剧"}, // Teenager
	{"年轻人:吃营养餐", "年轻人:喝水", "年轻人:看b站、看直播"}, // Younger
	{"老年人:吃饭", "老年人:喝脑x金", "老年人:看年代剧"},    // Older
}

const (
	eat     = 0
	drink   = 1
	watchTv = 2
)

type TableSims struct {
	StatusEnum StatusEnum
}

func NewTableSims() SimsController {
	return &TableSims{
		StatusEnum: Child,
	}
}

func (b *TableSims) Eat() {
	println(statusTable[b.StatusEnum][eat])
}
func (b *TableSims) Drink() {
	println(statusTable[b.StatusEnum][drink])
}
func (b *TableSims) WatchTv() {
	println(statusTable[b.StatusEnum][watchTv])
}
func (b *TableSims) Grow() {
	if b.StatusEnum < Older {
		b.StatusEnum++
	}
}

// 状态模式

type StatusController interface {
	setCurrentStatus(status status)
	Eat()
	Grow()
	Drink()
	WatchTv()
}

func NewStatusController() StatusController { return &Sims{currentStatus: &child{}} }

type Sims struct {
	currentStatus status
}

func (s *Sims) setCurrentStatus(status status) { s.currentStatus = status }
func (s *Sims) Eat()                           { s.currentStatus.Eat(s) }
func (s *Sims) Grow()                          { s.currentStatus.Grow(s) }
func (s *Sims) Drink()                         { s.currentStatus.Drink(s) }
func (s *Sims) WatchTv()                       { s.currentStatus.WatchTv(s) }

// 接口定义
type status interface {
	Eat(StatusController)
	Drink(StatusController)
	WatchTv(StatusController)
	Grow(StatusController)
}

// 子实现类
type child struct{}

func (b *child) Eat(controller StatusController)     { println("小孩子:喝稀饭") }
func (b *child) Drink(controller StatusController)   { println("小孩子:喝牛奶") }
func (b *child) WatchTv(controller StatusController) { println("小孩子:看卡通片") }
func (b *child) Grow(controller StatusController)    { controller.setCurrentStatus(new(teenager)) }

type teenager struct{}

func (b *teenager) Eat(controller StatusController)     { println("青少年:吃米线") }
func (b *teenager) Drink(controller StatusController)   { println("青少年:喝提神补脑液") }
func (b *teenager) WatchTv(controller StatusController) { println("青少年:看偶像剧") }
func (b *teenager) Grow(controller StatusController)    { controller.setCurrentStatus(new(younger)) }

type younger struct{}

func (b *younger) Eat(controller StatusController)     { println("年轻人:吃营养餐") }
func (b *younger) Drink(controller StatusController)   { println("年轻人:喝水") }
func (b *younger) WatchTv(controller StatusController) { println("年轻人:看b站、看直播") }
func (b *younger) Grow(controller StatusController)    { controller.setCurrentStatus(new(older)) }

type older struct{}

func (b *older) Eat(controller StatusController)     { println("老年人:吃饭") }
func (b *older) Drink(controller StatusController)   { println("老年人:喝脑x金") }
func (b *older) WatchTv(controller StatusController) { println("老年人:看年代剧") }
func (b *older) Grow(controller StatusController)    { /* Do nothing*/ }
