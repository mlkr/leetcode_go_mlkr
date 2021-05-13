package problem480

import "container/heap"

func medianSlidingWindow(nums []int, k int) []float64 {
	size := len(nums)
	res := make([]float64, 0, size-k+1)
	w := getWindow(nums, k)

	for i := k; i < size; i++ {
		median := w.getMedian(k)
		res = append(res, median)

		w.update(i-k, i, nums)
	}

	res = append(res, w.getMedian(k))

	return res
}

func getWindow(nums []int, k int) Window {
	var maxPQ MaxPQ
	var minPQ MinPQ
	m := make(map[int]*Entry)

	for i := 0; i < k; i++ {
		entry := &Entry{
			idx: i,
			val: nums[i],
		}

		m[i] = entry

		if len(minPQ) == len(maxPQ) {
			heap.Push(&minPQ, entry)
		} else {
			heap.Push(&maxPQ, entry)
		}
	}

	for len(maxPQ) > 0 && maxPQ[0].val > minPQ[0].val {
		maxPQ[0], minPQ[0] = minPQ[0], maxPQ[0]

		heap.Fix(&maxPQ, 0)
		heap.Fix(&minPQ, 0)
	}

	return Window{
		m:     m,
		maxPQ: maxPQ,
		minPQ: minPQ,
	}
}

type Entry struct {
	idx, val, index int
}

type Window struct {
	m     map[int]*Entry
	maxPQ MaxPQ
	minPQ MinPQ
}

func (w Window) getMedian(k int) float64 {
	if k&1 == 1 {
		return float64(w.minPQ[0].val)
	}

	return float64(w.minPQ[0].val+w.maxPQ[0].val) / 2
}

func (w Window) update(popId, pushId int, nums []int) {
	e := w.m[popId]
	e.idx = pushId
	e.val = nums[pushId]

	w.m[pushId] = e
	delete(w.m, popId)

	if e.index < len(w.minPQ) {
		heap.Fix(&w.minPQ, e.index)
	}

	if e.index < len(w.maxPQ) {
		heap.Fix(&w.maxPQ, e.index)
	}

	for len(w.maxPQ) > 0 && w.maxPQ[0].val > w.minPQ[0].val {
		w.maxPQ[0], w.minPQ[0] = w.minPQ[0], w.maxPQ[0]

		heap.Fix(&w.maxPQ, 0)
		heap.Fix(&w.minPQ, 0)
	}
}

type MaxPQ []*Entry

func (maxPQ MaxPQ) Len() int { return len(maxPQ) }

func (maxPQ MaxPQ) Less(i, j int) bool {
	return maxPQ[i].val > maxPQ[j].val
}

func (maxPQ MaxPQ) Swap(i, j int) {
	maxPQ[i], maxPQ[j] = maxPQ[j], maxPQ[i]
	maxPQ[i].index = i
	maxPQ[j].index = j
}

func (maxPQ *MaxPQ) Push(x interface{}) {
	n := len(*maxPQ)
	item := x.(*Entry)
	item.index = n
	*maxPQ = append(*maxPQ, item)
}

func (maxPQ *MaxPQ) Pop() interface{} {
	old := *maxPQ
	n := len(old)
	item := old[n-1]
	item.index = -1
	*maxPQ = old[0 : n-1]

	return item
}

type MinPQ []*Entry

func (minPQ MinPQ) Len() int { return len(minPQ) }

func (minPQ MinPQ) Less(i, j int) bool {
	return minPQ[i].val < minPQ[j].val
}

func (minPQ MinPQ) Swap(i, j int) {
	minPQ[i], minPQ[j] = minPQ[j], minPQ[i]
	minPQ[i].index = i
	minPQ[j].index = j
}

func (minPQ *MinPQ) Push(x interface{}) {
	n := len(*minPQ)
	item := x.(*Entry)
	item.index = n
	*minPQ = append(*minPQ, item)
}

func (minPQ *MinPQ) Pop() interface{} {
	old := *minPQ
	n := len(old)
	item := old[n-1]
	item.index = -1
	*minPQ = old[0 : n-1]

	return item
}
