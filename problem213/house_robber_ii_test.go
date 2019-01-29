package problem213

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{2, 3, 2},
		3,
	},

	{
		[]int{1, 2, 3, 1},
		4,
	},
}

func Test_rob(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(rob(q.nums), q.ans)
	}
}

func Benchmark_rob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rob(q.nums)
		}
	}
}
