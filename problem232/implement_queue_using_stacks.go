package problem232

type MyQueue struct {
	s1, s2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		s1: newStack(),
		s2: newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.Empty() {
		for !this.s1.Empty() {
			this.s2.Push(this.s1.Pop())
		}
	}

	return this.s2.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	x := this.Pop()
	this.s2.Push(x)
	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
type Stack struct {
	nums []int
}

func newStack() *Stack {
	return &Stack{
		nums: make([]int, 0),
	}
}

func (this *Stack) Push(x int) {
	this.nums = append(this.nums, x)
}

func (this *Stack) Pop() int {
	x := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return x
}

func (this *Stack) Empty() bool {
	return len(this.nums) == 0
}
