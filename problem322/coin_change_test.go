package problem322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	coins       []int
	amount, ans int
}{
	{
		[]int{1, 2, 5},
		11,
		3,
	},

	{
		[]int{2},
		3,
		-1,
	},
}

func Test_coinChange(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, coinChange(q.coins, q.amount), "输入 %v", q.coins)
	}
}

func Benchmark_coinChange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			coinChange(q.coins, q.amount)
		}
	}
}
