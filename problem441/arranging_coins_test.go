package problem441

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		5,
		2,
	},

	{
		6,
		3,
	},

	{
		8,
		3,
	},
}

func Test_arrangeCoins(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, arrangeCoins(q.n), "%d", q.n)
	}
}

func Test_arrangeCoins2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, arrangeCoins2(q.n), "%d", q.n)
	}
}

func Benchmark_arrangeCoins(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			arrangeCoins(q.n)
		}
	}
}

func Benchmark_arrangeCoins2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			arrangeCoins2(q.n)
		}
	}
}
