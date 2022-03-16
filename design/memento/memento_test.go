package memento

import "testing"

/* 备忘录模式
tip:
    场景:
        游戏存档、数据库binlog、redolog(增量备份)、redis备份

*/

func TestNewGameController(t *testing.T) {
    game := NewGameController("冬泳怪鸽")

    game.Beat()
    game.Info()
    game.StoreArchive()

    game.Beat()
    game.Info()
    game.LoadLastArchive()
    game.Info()
}
