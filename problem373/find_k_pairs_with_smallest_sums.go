package problem373

import "container/heap"

type Pair struct {
	sum   int
	n2Idx int
}

type Pairs []Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].sum < p[j].sum
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pairs) Push(x interface{}) {
	*p = append(*p, x.(Pair))
}

func (p *Pairs) Pop() interface{} {
	n := len(*p)
	old := *p

	v := (*p)[n-1]
	*p = old[:n-1]
	return v
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	pairs := make([]Pair, 0, len(nums1))
	for _, v := range nums1 {
		pairs = append(pairs, Pair{sum: v + nums2[0]})
	}

	heapPairs := Pairs(pairs)
	heap.Init(&heapPairs)

	if k > len(nums1)*len(nums2) {
		k = len(nums1) * len(nums2)
	}

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		v := heap.Pop(&heapPairs).(Pair)
		res[i] = []int{v.sum - nums2[v.n2Idx], nums2[v.n2Idx]}

		if v.n2Idx+1 < len(nums2) {
			heap.Push(
				&heapPairs,
				Pair{
					v.sum - nums2[v.n2Idx] + nums2[v.n2Idx+1],
					v.n2Idx + 1,
				},
			)
		}
	}

	return res
}
