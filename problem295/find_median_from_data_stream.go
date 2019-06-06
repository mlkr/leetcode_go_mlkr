package problem295

import "container/heap"

type MedianFinder struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{make([]int, 0, 128)}
}

func (this *MedianFinder) AddNum(num int) {
	size := len(this.nums)
	if size == 0 {
		this.nums = append(this.nums, num)
		return
	}

	left, right := 0, size-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if this.nums[mid] >= num {
			right = mid - 1
		}

		if this.nums[mid] < num {
			left = mid + 1
		}
	}

	if this.nums[mid] > num {
		mid--
	}

	temp := make([]int, size)
	copy(temp, this.nums)
	this.nums = append(append(temp[:mid+1], num), this.nums[mid+1:]...)
}

func (this *MedianFinder) FindMedian() float64 {
	size := len(this.nums)
	if size%2 == 1 {
		return float64(this.nums[size/2])
	}

	return float64((this.nums[size/2] + this.nums[size/2-1])) / 2
}

// 另一解法
type MedianFinder2 struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor2() MedianFinder2 {
	left := &MaxHeap{}
	heap.Init(left)

	right := &MinHeap{}
	heap.Init(right)
	return MedianFinder2{
		left:  left,
		right: right,
	}
}

func (this *MedianFinder2) AddNum2(num int) {
	if this.left.Len() == this.right.Len() {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	if this.right.Len() > 0 && this.left.IntHeap[0] > this.right.IntHeap[0] {
		this.left.IntHeap[0], this.right.IntHeap[0] = this.right.IntHeap[0], this.left.IntHeap[0]
		heap.Fix(this.left, 0)
		heap.Fix(this.right, 0)
	}
}

func (this *MedianFinder2) FindMedian2() float64 {
	res := float64(this.left.IntHeap[0])
	if this.left.Len() == this.right.Len() {
		res = (res + float64(this.right.IntHeap[0])) / 2
	}

	return res
}

type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less(i, j int) bool {
	return h.IntHeap[i] > h.IntHeap[j]
}

type MinHeap struct {
	IntHeap
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	size := h.Len()
	x := (*h)[size-1]
	*h = (*h)[:size-1]
	return x
}
