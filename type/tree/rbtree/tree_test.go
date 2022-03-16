package rbtree

import (
    "math/rand"
    "testing"
)

func TestNewRBTree(t *testing.T) {

    tree:=NewRBTree()



    for i:=0;i<10;i++{
        tree.Insert(rand.Uint64()%100000,rand.Uint64())
    }
    tree.LDR()



}