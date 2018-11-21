package problem155

type MinStack struct {
	stack []Stack
}

type Stack struct {
	min int
	num int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	min := x
	if len(this.stack) > 0 && this.GetMin() < x {
		min = this.GetMin()
	}

	this.stack = append(this.stack, Stack{num: x, min: min})
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1].num
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1].min
}
