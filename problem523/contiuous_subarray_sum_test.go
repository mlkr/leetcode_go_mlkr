package problem523

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  bool
}{

	{
		[]int{0},
		1,
		false,
	},

	{
		[]int{0, 0},
		1,
		true,
	},

	{
		[]int{2, 0, 0},
		6,
		true,
	},

	{
		[]int{2, 4},
		6,
		true,
	},

	{
		[]int{23, 2, 4, 6, 7},
		6,
		true,
	},

	{
		[]int{23, 2, 6, 4, 7},
		6,
		true,
	},

	{
		[]int{23, 2, 4, 6, 7},
		13,
		true,
	},

	{
		[]int{23, 2, 6, 4, 7},
		13,
		false,
	},
}

func Test_checkSubarraySum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, checkSubarraySum(q.nums, q.k), "输入 %v %v", q.nums, q.k)
	}
}

func Benchmark_checkSubarraySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkSubarraySum(q.nums, q.k)
		}
	}
}
