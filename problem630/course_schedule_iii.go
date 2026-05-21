package problem630

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) int {

	// 截止时间从小到大排序
	// 截止时间相同 持续时间从小到达排序
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return courses[i][1] < courses[j][1]
	})

	// 最大堆
	dh := &DurationHeap{}
	heap.Init(dh)
	duration := 0
	for _, c := range courses {
		heap.Push(dh, c[0])
		duration += c[0]

		for duration > c[1] {
			d := heap.Pop(dh)
			duration -= d.(int)
		}
	}

	return dh.Len()
}

type DurationHeap []int

func (dh DurationHeap) Len() int {
	return len(dh)
}

func (dh DurationHeap) Swap(i, j int) {
	dh[i], dh[j] = dh[j], dh[i]
}

func (dh DurationHeap) Less(i, j int) bool {
	return dh[i] > dh[j]
}

func (dh *DurationHeap) Push(x interface{}) {
	*dh = append(*dh, x.(int))
}
func (dh *DurationHeap) Pop() (x interface{}) {
	size := len(*dh)
	x = (*dh)[size-1]
	*dh = (*dh)[:size-1]

	return x
}
