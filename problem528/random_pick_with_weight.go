package problem528

import (
	"math/rand"
	"sort"
)

type Solution struct {
	total  int
	counts []int
}

func Constructor(w []int) Solution {
	s := Solution{}
	s.counts = make([]int, len(w))
	for i, times := range w {
		s.total += times
		s.counts[i] = s.total
	}

	return s
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.total)
	i := sort.Search(len(this.counts), func(i int) bool {
		return this.counts[i] > n
	})

	return i
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
