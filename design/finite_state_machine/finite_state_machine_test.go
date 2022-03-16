package finite_state_machine

import "testing"

/* 状态机
tip:
    场景:
        游戏、工作引擎
    desc:
        实现有三种、分支逻辑法(优点就是简单易实现,缺点就是耦合度、可读性、可维护性都很差。)、查表法(用二维数组来表示某一个状态的行为,复杂的状态机使用查表法就会非常困难)、状态模式(将逻辑和实现进行抽离,每一种状态都会有一个实现类)
    状态模式:状态模式(将逻辑和实现进行抽离,每一种状态都会有一个实现类)
*/

func TestFiniteStateMachine(t *testing.T) {
    t.Run("分支法", func(t *testing.T) {
        controller := NewBranchSims()
        for i := Child; i <= Older; i++ {
            controller.Eat()
            controller.WatchTv()
            controller.Drink()
            controller.Grow()
        }

    })

    t.Run("查表法", func(t *testing.T) {
        controller := NewTableSims()
        for i := Child; i <= Older; i++ {
            controller.Eat()
            controller.WatchTv()
            controller.Drink()
            controller.Grow()
        }
    })

    t.Run("状态模式", func(t *testing.T) {
        controller := NewStatusController()
        for i := Child; i <= Older; i++ {
            controller.Eat()
            controller.WatchTv()
            controller.Drink()
            controller.Grow()
        }
    })

}
