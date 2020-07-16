package problem416

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{5, 4, 3, 2, 2},
		true,
	},

	{
		[]int{3, 3, 3, 4, 5},
		true,
	},

	{
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100,
		},
		false,
	},

	{
		[]int{1, 2, 5},
		false,
	},

	{
		[]int{1, 5, 11, 5},
		true,
	},

	{
		[]int{1, 2, 3, 5},
		false,
	},
}

func Test_canPartition(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, canPartition(q.nums), "输入 %v", q.nums)
	}
}

func Test_canPartition2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, canPartition2(q.nums), "输入 %v", q.nums)
	}
}

func Benchmark_canPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canPartition(q.nums)
		}
	}
}

func Benchmark_canPartition2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canPartition2(q.nums)
		}
	}
}
