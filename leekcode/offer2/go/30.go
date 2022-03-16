package _go

import "math/rand"


type RandomizedSet struct {
	m   map[int]int
	val []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m:   make(map[int]int),
		val: []int{},
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.val = append(this.val, val)
		this.m[val] = len(this.val) - 1
		return true
	}

	return false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {

	idx, ok := this.m[val]
	if !ok {
		return false
	}

	delete(this.m, val)
	num := this.val[len(this.val)-1]
	this.val = this.val[:len(this.val)-1]

	if num != val {
		this.val[idx] = num
		this.m[num] = idx
	}
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.val))
	return this.val[index]
}