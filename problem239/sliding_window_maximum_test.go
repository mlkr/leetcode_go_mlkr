package problem239

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  []int
}{
	{
		[]int{},
		0,
		[]int{},
	},

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]int{3, 3, 5, 5, 6, 7},
	},

	{
		[]int{7, 2, 4},
		2,
		[]int{7, 4},
	},
}

func Test_maxSlidingWindow(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxSlidingWindow(q.nums, q.k))
	}
}

func Test_maxSlidingWindow2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxSlidingWindow2(q.nums, q.k))
	}
}

func Benchmark_maxSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxSlidingWindow(q.nums, q.k)
		}
	}
}

func Benchmark_maxSlidingWindow2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxSlidingWindow2(q.nums, q.k)
		}
	}
}
