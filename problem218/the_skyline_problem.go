package problem218

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	h := new(heightHeap)
	heap.Init(h)
	pre := 0
	heap.Push(h, pre)

	edges := make([][3]int, 0, len(buildings)*2)
	for _, b := range buildings {
		// 进入边
		edges = append(edges, [3]int{b[0], b[2], -1})
		// 退出边
		edges = append(edges, [3]int{b[1], b[2], 1})

	}
	sort.Sort(edgeSort(edges))

	points := [][]int{}
	for _, e := range edges {
		if e[2] < 0 {
			// 进入边, 加入高度
			heap.Push(h, e[1])
		} else {
			// 退出边, 去掉高度
			i := 0
			for i < len(*h) {
				if (*h)[i] == e[1] {
					break
				}
				i++
			}

			heap.Remove(h, i)
		}

		now := (*h)[0]
		if now != pre {
			points = append(points, []int{e[0], now})
			pre = now
		}
	}

	return points
}

type heightHeap []int

func (h heightHeap) Len() int           { return len(h) }
func (h heightHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h heightHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *heightHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *heightHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type edgeSort [][3]int

func (e edgeSort) Len() int { return len(e) }
func (e edgeSort) Less(i, j int) bool {
	if e[i][0] != e[j][0] {
		return e[i][0] < e[j][0]
	}

	return e[i][1]*e[i][2] < e[j][1]*e[j][2]
}
func (e edgeSort) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
