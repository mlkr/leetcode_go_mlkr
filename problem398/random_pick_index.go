package problem398

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Pick(target int) int {
	count := 0
	res := -1
	for k, v := range this.nums {
		if v != target {
			continue
		}

		count++
		if rand.Intn(count) == 0 {
			res = k
		}
	}

	return res
}
