package problem309

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	prices []int
	profit int
}{
	{
		[]int{1, 2, 3, 0, 2},
		3,
	},

	{
		[]int{2, 1, 4},
		3,
	},

	{
		[]int{1, 7, 2, 4},
		6,
	},

	{
		[]int{6, 1, 6, 4, 3, 0, 2},
		7,
	},

	{
		[]int{6},
		0,
	},
}

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.profit, maxProfit(q.prices))
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxProfit(q.prices)
		}
	}
}
