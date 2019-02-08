package problem215

import (
	"container/heap"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	quickSort(nums)
	return nums[len(nums)-k]
}

func quickSort(nums []int) {
	size := len(nums)
	if size <= 1 {
		return
	}

	point := part(nums)
	quickSort(nums[:point])
	quickSort(nums[point+1:])

}

func part(nums []int) int {
	size := len(nums)
	n := rand.Intn(size)
	nums[0], nums[n] = nums[n], nums[0]
	i, j := 1, size-1
	point := 0

	for i < j {
		for i < size-1 && nums[i] <= nums[0] {
			i++
		}

		for j > 1 && nums[0] < nums[j] {
			j--
		}

		if i < j && nums[i] > nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if nums[0] >= nums[j] {
		nums[0], nums[j] = nums[j], nums[0]
		point = j
	}

	return point
}

// 解法二
func findKthLargest2(nums []int, k int) int {
	temp := IntHeap(nums)
	h := &temp
	heap.Init(h)

	if k == 1 {
		return (*h)[0]
	}

	for i := 1; i < k; i++ {
		heap.Remove(h, 0)
	}

	return (*h)[0]
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

// 从大到小
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return res
}
