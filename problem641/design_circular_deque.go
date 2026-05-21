package problem641

type MyCircularDeque struct {
	num   []int
	count int
	f, r  int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		num: make([]int, k),
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.f = 1
		this.r = 0
	}

	size := len(this.num)
	this.f = (this.f - 1 + size) % size
	this.num[this.f] = value

	this.count++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	if this.IsEmpty() {
		this.r = -1
		this.f = 0
	}

	size := len(this.num)
	this.r = (this.r + 1) % size
	this.num[this.r] = value

	this.count++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	size := len(this.num)
	this.f = (this.f + 1) % size

	this.count--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	size := len(this.num)
	this.r = (this.r - 1 + size) % size

	this.count--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.num[this.f]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.num[this.r]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.count == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.count == len(this.num)
}
