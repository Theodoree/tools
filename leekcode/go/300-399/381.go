package _00_399

import "fmt"

/*
381. O(1) 时间插入、删除和获取随机元素 - 允许重复
*/

type RandomizedCollection struct {
	data map[string]int
	key  map[int][]string
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		data: make(map[string]int),
		key:  make(map[int][]string),
	}

}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {

	if v, ok := this.key[val]; ok {
		this.key[val] = append(v, fmt.Sprintf("%d_%d", val, len(v)))
	}else{
		this.key[val] = []string{ fmt.Sprintf("%d_%d", val, 0)}
	}


	return true
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {

}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {

}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
