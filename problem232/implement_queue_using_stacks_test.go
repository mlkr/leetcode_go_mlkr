package problem232

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyQueue(t *testing.T) {
	ast := assert.New(t)
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	ast.Equal(1, queue.Peek())
	ast.Equal(1, queue.Pop())
	ast.False(queue.Empty())
}

func Benchmark_MyQueue(b *testing.B) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Peek()
	queue.Pop()
	queue.Empty()
}
