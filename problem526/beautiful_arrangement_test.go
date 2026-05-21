package problem526

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		2,
		2,
	},

	{
		1,
		1,
	},
}

func Test_countArrangement(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, countArrangement(q.n), "输入 %v", q.n)
	}
}

func Benchmark_countArrangement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countArrangement(q.n)
		}
	}
}
