package problem532

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums   []int
	k, ans int
}{
	{
		[]int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3},
		3,
		2,
	},

	{
		[]int{3, 1, 4, 1, 5},
		2,
		2,
	},

	{
		[]int{1, 2, 3, 4, 5},
		1,
		4,
	},

	{
		[]int{1, 3, 1, 5, 4},
		0,
		1,
	},
}

func Test_findPairs(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findPairs(q.nums, q.k), "输入 %v", q.nums)
	}
}

func Benchmark_findPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findPairs(q.nums, q.k)
		}
	}
}

func Test_findPairs2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findPairs2(q.nums, q.k), "输入 %v", q.nums)
	}
}

func Benchmark_findPairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findPairs2(q.nums, q.k)
		}
	}
}
