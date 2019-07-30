package problem312

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  int
}{
	{
		[]int{3, 1, 5, 8},
		167,
	},

	{
		[]int{},
		0,
	},
}

func Test_maxCoins(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxCoins(q.nums))
	}
}

func Benchmark_maxCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxCoins(q.nums)
		}
	}
}
