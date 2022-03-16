package worker

import (
    "fmt"
    "testing"
)

func BenchmarkNewWorkPool(b *testing.B) {
    b.ReportAllocs()
    pool, err := NewWorkPool(10, 100)
    if err != nil {
        b.Fatal(err)
    }

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            pool.PushTask(test1, panicHandler)
        }
    })
    pool.Close()

}

func BenchmarkNormalHandler(b *testing.B) {
    b.ReportAllocs()
    var c = make(chan struct{}, 10)

    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            c <- struct{}{}
            go func() {
                defer func() {
                    <- c
                }()
                test1()
            }()
        }
    })
}

func test1() {
    fib(12)
}

func panicHandler(err error) {
    fmt.Printf("%v \n", err)
}

// 简单粗暴,直接来递归fib函数
func fib(i int) int {
    if i == 0 {
        return 0
    }
    if i == 1 {
        return 1
    }
    return fib(i-1) + fib(i-2)
}
