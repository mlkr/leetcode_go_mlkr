package problem303

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumArray(t *testing.T) {
	ast := assert.New(t)
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	ast.Equal(1, numArray.SumRange(0, 2))
	ast.Equal(-1, numArray.SumRange(2, 5))
	ast.Equal(-3, numArray.SumRange(0, 5))
}

func Benchmark_NumArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
		numArray.SumRange(0, 2)
		numArray.SumRange(2, 5)
		numArray.SumRange(0, 5)
	}
}
