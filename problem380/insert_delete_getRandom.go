package problem380

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	m    map[int]int
	nums []int
}

func init() {
	rand.Seed(time.Now().Unix())
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    make(map[int]int, 0),
		nums: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	nIdx, ok := this.m[val]
	if !ok {
		return false
	}

	size := len(this.nums)
	if nIdx != size-1 {
		this.nums[nIdx] = this.nums[size-1]
		this.m[this.nums[size-1]] = nIdx
	}
	this.nums = this.nums[:size-1]

	delete(this.m, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
