package bptree

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestNewBpTree(t *testing.T) {

	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("%.2fM object %d \n", float64(stats.HeapAlloc)/1024/1024, stats.HeapObjects)
	tree := NewBpTree()
	rand.Seed(time.Now().Unix())

	for i := 1; i < 100000; i++ {
		tree.Insert(uint64(i), uint64(i))
	}

	tree.findAndPrint(150)

	tree.Delete(150)

	tree.findAndPrint(150)

	runtime.GC()
	runtime.ReadMemStats(&stats)
	fmt.Printf("%.2fM object %d \n", float64(stats.HeapAlloc)/1024/1024, stats.HeapObjects)

	tree.DestroyTree()

	runtime.GC()
	runtime.ReadMemStats(&stats)
	fmt.Printf("%.2fM object %d \n", float64(stats.HeapAlloc)/1024/1024, stats.HeapObjects)

}

func BenchmarkNewBpTree(b *testing.B) {
	rand.Seed(time.Now().Unix())

}
