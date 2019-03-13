package problem225

type MyStack struct {
	a, b *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		a: NewQueue(),
		b: NewQueue(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.a.Empty() {
		this.a, this.b = this.b, this.a
	}

	this.a.Push(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.a.Empty() {
		this.a, this.b = this.b, this.a
	}
	for this.a.Len() > 1 {
		this.b.Push(this.a.Pop())
	}
	return this.a.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	num := this.Pop()
	this.Push(num)
	return num
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.a.Empty() && this.b.Empty()
}

// 队列
type Queue struct {
	nums []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(x int) {
	q.nums = append(q.nums, x)
}

func (q *Queue) Pop() int {
	num := q.nums[0]
	q.nums = q.nums[1:]
	return num
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) Empty() bool {
	return len(q.nums) == 0
}
