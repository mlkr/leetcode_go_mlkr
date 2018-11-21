package problem155

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minStack(t *testing.T) {
	ast := assert.New(t)

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	ast.Equal(minStack.GetMin(), -3)
	minStack.Pop()
	ast.Equal(minStack.Top(), 0)
	ast.Equal(minStack.GetMin(), -2)
}

func Benchmark_minStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		minStack.GetMin()
		minStack.Pop()
		minStack.Top()
		minStack.GetMin()
	}
}
