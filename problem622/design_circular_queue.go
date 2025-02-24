package problem622

type MyCircularQueue struct {
	l, r  int
	count int
	nums  []int
}

func Constructor(k int) MyCircularQueue {
	nums := make([]int, k)
	for i := range nums {
		nums[i] = -1
	}

	mc := MyCircularQueue{0, 0, 0, nums}
	return mc
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	i := this.r
	if this.nums[i] != -1 {
		return false
	}

	this.nums[i] = value
	this.r = (i + 1) % len(this.nums)
	this.count++
	return true
}

// 删除队首元素
func (this *MyCircularQueue) DeQueue() bool {
	i := this.l
	if this.nums[i] == -1 {
		return false
	}

	this.nums[i] = -1
	i++
	this.l = i % len(this.nums)
	this.count--

	return true
}

func (this *MyCircularQueue) Front() int {
	return this.nums[this.l]
}

func (this *MyCircularQueue) Rear() int {
	i := this.r - 1
	if i == -1 {
		i = len(this.nums) - 1
	}
	return this.nums[i]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.count == len(this.nums)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
