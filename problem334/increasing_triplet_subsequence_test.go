package problem334

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  bool
}{
	{
		[]int{1, 5, 4, 2, 3},
		true,
	},

	{
		[]int{3, 4, 1, 5},
		true,
	},

	{
		[]int{5, 1, 5, 5, 2, 5, 4},
		true,
	},

	{
		[]int{1, 2, 3, 1, 2, 1},
		true,
	},

	{
		[]int{2, 1, 5, 0, 4, 6},
		true,
	},

	{
		[]int{1, 2, 3, 4, 5},
		true,
	},

	{
		[]int{5, 4, 3, 2, 1},
		false,
	},
}

func Test_increasingTriplet(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, increasingTriplet(q.nums), "输入%v", q.nums)
	}
}

func Benchmark_increasingTriplet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			increasingTriplet(q.nums)
		}
	}
}
