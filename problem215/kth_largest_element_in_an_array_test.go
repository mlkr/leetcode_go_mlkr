package problem215

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input  []int
	k      int
	output int
}{
	{
		[]int{
			3, 2, 1, 5, 6, 4,
		},
		2,
		5,
	},

	{
		[]int{
			3, 2, 3, 1, 2, 4, 5, 5, 6,
		},
		4,
		4,
	},
}

func Test_findKthLargest(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(findKthLargest(q.input, q.k), q.output)
	}
}

func Test_findKthLargest2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(findKthLargest2(q.input, q.k), q.output)
	}
}
func Benchmark_findKthLargest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findKthLargest(q.input, q.k)
		}
	}
}

func Benchmark_findKthLargest2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findKthLargest2(q.input, q.k)
		}
	}
}

func Test_quickSort(t *testing.T) {
	questions := []struct {
		input, output []int
	}{
		{
			[]int{0},
			[]int{0},
		},

		{
			[]int{1, 1},
			[]int{1, 1},
		},

		{
			[]int{1, 0},
			[]int{0, 1},
		},

		{
			[]int{0, 1},
			[]int{0, 1},
		},

		{
			[]int{3, 4, 2},
			[]int{2, 3, 4},
		},

		{
			[]int{2, 5, 8, 2, 6, 7, 3},
			[]int{2, 2, 3, 5, 6, 7, 8},
		},
	}

	ast := assert.New(t)
	for _, q := range questions {
		quickSort(q.input)
		ast.Equal(q.input, q.output)
	}
}
