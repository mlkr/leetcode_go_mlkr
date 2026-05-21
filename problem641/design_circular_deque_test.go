package problem641

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyCircularDeque(t *testing.T) {
	ast := assert.New(t)
	// mcd := Constructor(3)
	// ast.Equal(-1, mcd.GetFront())
	// ast.Equal(-1, mcd.GetRear())

	// ast.Equal(true, mcd.InsertLast(1))
	// ast.Equal(true, mcd.InsertLast(2))
	// ast.Equal(true, mcd.InsertFront(3))
	// ast.Equal(false, mcd.InsertFront(4))
	// ast.Equal(false, mcd.InsertLast(4))
	// ast.Equal(2, mcd.GetRear())
	// ast.Equal(true, mcd.IsFull())
	// ast.Equal(true, mcd.DeleteLast())
	// ast.Equal(true, mcd.InsertFront(4))
	// ast.Equal(4, mcd.GetFront())
	// ast.Equal(true, mcd.DeleteFront())
	// ast.Equal(3, mcd.GetFront())

	mcd := Constructor(2)
	ast.Equal(true, mcd.InsertFront(7))
	ast.Equal(true, mcd.DeleteLast())
	ast.Equal(-1, mcd.GetFront())
	ast.Equal(true, mcd.InsertLast(5))
	ast.Equal(true, mcd.InsertFront(0))
	ast.Equal(0, mcd.GetFront())
	ast.Equal(5, mcd.GetRear())
	ast.Equal(0, mcd.GetFront())
	ast.Equal(0, mcd.GetFront())
	ast.Equal(5, mcd.GetRear())
	ast.Equal(false, mcd.InsertLast(0))
}

func Benchmark_MyCircularDeque(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
