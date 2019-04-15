package problem240

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix [][]int
	target int
	ans    bool
}{
	{
		[][]int{},
		5,
		false,
	},

	{
		[][]int{
			[]int{},
		},
		5,
		false,
	},

	{
		[][]int{
			[]int{1, 4, 7, 11, 15},
			[]int{2, 5, 8, 12, 19},
			[]int{3, 6, 9, 16, 22},
			[]int{10, 13, 14, 17, 24},
			[]int{18, 21, 23, 26, 30},
		},
		5,
		true,
	},

	{
		[][]int{
			[]int{1, 4, 7, 11, 15},
			[]int{2, 5, 8, 12, 19},
			[]int{3, 6, 9, 16, 22},
			[]int{10, 13, 14, 17, 24},
			[]int{18, 21, 23, 26, 30},
		},
		20,
		false,
	},

	{
		[][]int{
			[]int{1, 4, 7, 11, 15},
			[]int{2, 5, 8, 12, 19},
			[]int{3, 6, 9, 16, 22},
			[]int{10, 13, 14, 17, 24},
			[]int{18, 21, 23, 26, 30},
		},
		0,
		false,
	},
}

func Test_searchMatrix(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(searchMatrix(q.matrix, q.target), q.ans)
	}
}

func Test_searchMatrix2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(searchMatrix2(q.matrix, q.target), q.ans)
	}
}

func Benchmark_searchMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			searchMatrix(q.matrix, q.target)
		}
	}
}

func Benchmark_searchMatrix2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			searchMatrix2(q.matrix, q.target)
		}
	}
}
