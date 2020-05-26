package problem381

import (
	"math/rand"
	"sort"
	"time"
)

type RandomizedCollection struct {
	m map[int][]int
	nums []int
}

func init() {
	rand.Seed(time.Now().Unix())
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m: make(map[int][]int, 0),
		nums: make([]int, 0),
	}
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.nums = append(this.nums, val)
	nIdx := len(this.nums)-1

	nIdxs, ok := this.m[val]
	if !ok {
		this.m[val] = []int{nIdx}
		return true
	}

	nIdxs = append(nIdxs, nIdx)
	this.m[val] = nIdxs
	return false
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
    nIdxs, ok := this.m[val]
	if !ok {
		return false
	}

	nIdx := nIdxs[len(nIdxs)-1]
	size := len(this.nums)
	if nIdx != size-1 {
		this.nums[nIdx] = this.nums[size-1]
		lastNidxs := this.m[this.nums[size-1]]
		lastNidxs[len(lastNidxs)-1] = nIdx
		sort.Ints(this.m[this.nums[size-1]])
	}
	this.nums = this.nums[:size-1]

	if len(nIdxs) != 1{
		this.m[val] = nIdxs[:len(nIdxs)-1]
	}else{
		delete(this.m, val)
	}

	return true
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]    
}