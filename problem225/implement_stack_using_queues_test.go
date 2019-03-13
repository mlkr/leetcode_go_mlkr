package problem225

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	ast := assert.New(t)

	q := NewQueue()
	ast.True(q.Empty())

	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	ast.Equal(100, q.Len())

	for i := 0; i < 100; i++ {
		ast.Equal(i, q.Pop())
	}
}

func Benchmark_Queue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue()
		q.Empty()

		for i := 0; i < 100; i++ {
			q.Push(i)
		}

		for i := 0; i < 100; i++ {
			q.Pop()
		}
	}
}

func Test_Mystack(t *testing.T) {
	ast := assert.New(t)
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	ast.Equal(2, stack.Top())
	ast.Equal(2, stack.Pop())
	ast.False(stack.Empty())
	ast.Equal(1, stack.Pop())
	ast.True(stack.Empty())
}

func Benchmark_Mystack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		stack.Top()
		stack.Pop()
		stack.Empty()
	}
}
