package problem188

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	prices []int
	k      int
	ans    int
}{
	{
		[]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
		2,
		13,
	},

	{
		[]int{2, 4, 1},
		2,
		2,
	},

	{
		[]int{3, 2, 6, 5, 0, 3},
		2,
		7,
	},
}

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(maxProfit(q.k, q.prices), q.ans)
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxProfit(q.k, q.prices)
		}
	}
}
