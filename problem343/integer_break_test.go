package problem343

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n   int
	ans int
}{
	{
		3,
		2,
	},

	{
		2,
		1,
	},

	{
		10,
		36,
	},
}

func Test_integerBreak(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, integerBreak(q.n))
	}
}

func Benchmark_integerBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			integerBreak(q.n)
		}
	}
}
