package memento

import (
    "encoding/json"
    "fmt"
)

type GameController interface {
    StoreArchive()
    Beat()
    LoadLastArchive()
    Info()
}

type CheckPoint uint64

const (
    village CheckPoint = iota + 1
    grassland
    castle
)

type gameController struct {
    Data   *gameData
    Holder snapshotHolder
}
type gameData struct {
    Name       string     `json:"name"`
    Checkpoint CheckPoint `json:"checkpoint"`
    Hp         uint64     `json:"hp"`
    Mp         uint64     `json:"mp"`
}

func NewGameController(name string) GameController {
    return &gameController{
        Data: &gameData{
            Name:       name,
            Checkpoint: village,
            Hp:         100,
            Mp:         100,
        },
        Holder: newSnapshotHolder(),
    }
}
func (controller *gameController) Beat() {
    if controller.Data.Checkpoint <= castle {
        controller.Data.Checkpoint++
    }
    println("欧拉欧拉欧拉欧拉")
    controller.Data.Hp += 100
    controller.Data.Mp += 100
    println("已通关;当前关卡:", controller.Data.Checkpoint)
}
func (controller *gameController) StoreArchive() {
    snapShot := NewSnapShot(controller.Data)
    controller.Holder.Push(snapShot)
}
func (controller *gameController) LoadLastArchive() {
    snapshot := controller.Holder.Pop()
    if snapshot == nil {
        println("还未有存档")
    } else {
        println("已加载存档")
        json.Unmarshal(snapshot.GetContext(), controller.Data)
    }

}

func (controller *gameController) Info() {
    fmt.Printf("%#v \n", controller.Data)
}

type SnapShot interface {
    GetContext() []byte
}
type snapshot struct {
    context []byte
}

func NewSnapShot(game *gameData) *snapshot {
    buf, err := json.Marshal(game)
    if err != nil {
        panic(err)
    }
    return &snapshot{context: buf}
}
func (s *snapshot) GetContext() []byte {
    return s.context
}

type snapshotHolder interface {
    Push(*snapshot)
    Pop() *snapshot
}
type SnapshotHolder struct {
    arr []*snapshot
}

func newSnapshotHolder() snapshotHolder {
    return &SnapshotHolder{}
}
func (s *SnapshotHolder) Push(snapshot *snapshot) {
    s.arr = append(s.arr, snapshot)
}
func (s *SnapshotHolder) Pop() *snapshot {
    if len(s.arr) > 0 {
        return s.arr[0]
    }
    return nil
}
