package avltree

import (
    "fmt"
    "testing"
)

func TestAVLTreePut(t *testing.T) {
    tree := NewWithIntComparator()
    tree.Put(5, "e")
    tree.Put(6, "f")
    tree.Put(7, "g")
    tree.Put(3, "c")
    tree.Put(4, "d")
    tree.Put(1, "x")
    tree.Put(2, "b")
    tree.Put(1, "a") //overwrite


    if actualValue := tree.Size(); actualValue != 7 {
        t.Errorf("Got %v expected %v", actualValue, 7)
    }
    if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d%d%d%d", tree.Keys()...), "1234567"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s%s%s%s", tree.Values()...), "abcdefg"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }

    tests1 := [][]interface{}{
        {1, "a", true},
        {2, "b", true},
        {3, "c", true},
        {4, "d", true},
        {5, "e", true},
        {6, "f", true},
        {7, "g", true},
        {8, nil, false},
    }

    for _, test := range tests1 {
        // retrievals
        actualValue, actualFound := tree.Get(test[0])
        if actualValue != test[1] || actualFound != test[2] {
            t.Errorf("Got %v expected %v", actualValue, test[1])
        }
    }
}

func TestAVLTreeRemove(t *testing.T) {
    tree := NewWithIntComparator()
    tree.Put(5, "e")
    tree.Put(6, "f")
    tree.Put(7, "g")
    tree.Put(3, "c")
    tree.Put(4, "d")
    tree.Put(1, "x")
    tree.Put(2, "b")
    tree.Put(1, "a") //overwrite

    tree.Delete(5)
    tree.Delete(6)
    tree.Delete(7)
    tree.Delete(8)
    tree.Delete(5)

    if actualValue, expectedValue := fmt.Sprintf("%d%d%d%d", tree.Keys()...), "1234"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if actualValue, expectedValue := fmt.Sprintf("%s%s%s%s", tree.Values()...), "abcd"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if actualValue := tree.Size(); actualValue != 4 {
        t.Errorf("Got %v expected %v", actualValue, 7)
    }

    tests2 := [][]interface{}{
        {1, "a", true},
        {2, "b", true},
        {3, "c", true},
        {4, "d", true},
        {5, nil, false},
        {6, nil, false},
        {7, nil, false},
        {8, nil, false},
    }

    for _, test := range tests2 {
        actualValue, actualFound := tree.Get(test[0])
        if actualValue != test[1] || actualFound != test[2] {
            t.Errorf("Got %v expected %v", actualValue, test[1])
        }
    }

    tree.Delete(1)
    tree.Delete(4)
    tree.Delete(2)
    tree.Delete(3)
    tree.Delete(2)
    tree.Delete(2)

    if actualValue, expectedValue := fmt.Sprintf("%s", tree.Keys()), "[]"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if actualValue, expectedValue := fmt.Sprintf("%s", tree.Values()), "[]"; actualValue != expectedValue {
        t.Errorf("Got %v expected %v", actualValue, expectedValue)
    }
    if empty, size := tree.Empty(), tree.Size(); empty != true || size != -0 {
        t.Errorf("Got %v expected %v", empty, true)
    }

}
