package problem643

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  float64
}{
	{
		[]int{-1},
		1,
		-1.0,
	},

	{
		[]int{1, 12, -5, -6, 50, 3},
		4,
		12.75,
	},

	{
		[]int{5},
		1,
		5.0,
	},
}

func Test_findMaxAverage(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findMaxAverage(q.nums, q.k))
	}
}

func Benchmark_findMaxAverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMaxAverage(q.nums, q.k)
		}
	}
}
