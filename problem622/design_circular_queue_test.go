package problem622

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCircularQueue(t *testing.T) {
	ast := assert.New(t)

	// mc := Constructor(3)
	// ast.True(mc.EnQueue(1))
	// ast.True(mc.EnQueue(2))
	// ast.True(mc.EnQueue(3))
	// ast.False(mc.EnQueue(4))

	// ast.Equal(3, mc.Rear())
	// ast.True(mc.IsFull())

	// ast.True(mc.DeQueue())
	// ast.True(mc.EnQueue(4))
	// ast.Equal(4, mc.Rear())

	mc := Constructor(8)
	ast.True(mc.EnQueue(3))
	ast.True(mc.EnQueue(9))
	ast.True(mc.EnQueue(5))
	ast.True(mc.EnQueue(0))

	ast.True(mc.DeQueue())
	ast.True(mc.DeQueue())

	ast.False(mc.IsEmpty())
	ast.False(mc.IsEmpty())

	ast.Equal(0, mc.Rear())
	ast.Equal(0, mc.Rear())

	ast.True(mc.DeQueue())
}
