package problem641

type MyCircularDeque struct {
	nums        []int
	size        int
	front, last int
}

func Constructor(k int) MyCircularDeque {
	mcd := MyCircularDeque{
		nums: make([]int, k),
	}

	return mcd
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	n := len(this.nums)
	if !this.IsEmpty() {
		this.front = (this.front + n - 1) % n
	}

	this.nums[this.front] = value
	this.size++

	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	n := len(this.nums)
	if !this.IsEmpty() {
		this.last = (this.last + 1) % n
	}

	this.nums[this.last] = value
	this.size++

	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.size--
	if this.IsEmpty() {
		this.front, this.last = 0, 0
		return true
	}

	n := len(this.nums)
	this.front = (this.front + 1) % n

	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.size--
	if this.IsEmpty() {
		this.front, this.last = 0, 0
		return true
	}

	n := len(this.nums)
	this.last = (this.last + n - 1) % n

	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.nums[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.nums[this.last]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.nums)
}
