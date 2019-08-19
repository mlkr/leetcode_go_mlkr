package problem315

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  []int
}{
	{
		[]int{5, 2, 6, 1},
		[]int{2, 1, 1, 0},
	},

	{
		[]int{1, 5, 2, 6, 1},
		[]int{0, 2, 1, 1, 0},
	},

	{
		[]int{5, 2, 6, 1, 1},
		[]int{3, 2, 2, 0, 0},
	},

	{
		[]int{},
		[]int{},
	},
}

func Test_countSmaller(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, countSmaller(q.nums))
	}
}

func Benchmark_countSmaller(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countSmaller(q.nums)
		}
	}
}
