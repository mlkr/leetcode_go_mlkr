package problem650

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		12,
		7,
	},
	// 111111 111111
	// sub +2 =
	// 3+2 +2 = 7

	// 111 111 111 111
	// 3+4 = 7

	// 11 11 11 11 11 11
	// 2+6 = 8

	// 111 111
	// 3+2 = 6
	// 11 11 11
	// 2+3 = 6

	{
		3,
		3,
	},

	{
		1,
		0,
	},
}

func Test_minSteps(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, minSteps(q.n), "输入 %v", q.n)
	}
}

func Benchmark_minSteps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minSteps(q.n)
		}
	}
}
