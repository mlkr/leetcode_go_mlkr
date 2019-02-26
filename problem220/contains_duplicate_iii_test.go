package problem220

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k, t int
	ans  bool
}{
	{
		[]int{-1, -1},
		1,
		-1,
		false,
	},

	{
		[]int{1, 2, 3, 1},
		3,
		0,
		true,
	},

	{
		[]int{1, 0, 1, 1},
		1,
		2,
		true,
	},

	{
		[]int{1, 5, 9, 1, 5, 9},
		2,
		3,
		false,
	},
}

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, containsNearbyAlmostDuplicate(q.nums, q.k, q.t), "输入 %v", q.nums)
	}
}

func Benchmark_containsNearbyAlmostDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			containsNearbyAlmostDuplicate(q.nums, q.k, q.t)
		}
	}
}
