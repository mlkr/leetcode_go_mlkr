package problem304

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var matrix = [][]int{
	[]int{3, 0, 1, 4, 2},
	[]int{5, 6, 3, 2, 1},
	[]int{1, 2, 0, 1, 5},
	[]int{4, 1, 0, 1, 7},
	[]int{1, 0, 3, 0, 5},
}

var questions = []struct {
	row1, col1, row2, col2, ans int
}{
	{2, 1, 4, 3, 8},
	{1, 1, 2, 2, 11},
	{1, 2, 2, 4, 12},
}

var matrix2 = [][]int{}

var questions2 = []struct {
	row1, col1, row2, col2, ans int
}{
	{0, 0, 0, 0, 0},
}

func Test_NumMatrix(t *testing.T) {
	ast := assert.New(t)
	numMatrix := Constructor(matrix)
	for _, q := range questions {
		ast.Equal(q.ans, numMatrix.SumRegion(q.row1, q.col1, q.row2, q.col2))
	}

	numMatrix2 := Constructor(matrix2)
	for _, q := range questions2 {
		ast.Equal(q.ans, numMatrix2.SumRegion(q.row1, q.col1, q.row2, q.col2))
	}
}

func Benchmark_NumMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numMatrix := Constructor(matrix)
		for _, q := range questions {
			numMatrix.SumRegion(q.row1, q.col1, q.row2, q.col2)
		}
	}
}
