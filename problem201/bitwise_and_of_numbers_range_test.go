package problem201

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	m, n int
	ans  int
}{
	// 0101
	// 0110
	// 0111
	{
		5, 7,
		4,
	},

	{
		0, 1,
		0,
	},
}

func Test_rangeBitwiseAnd(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(rangeBitwiseAnd(q.m, q.n), q.ans)
	}
}

func Benchmark_rangeBitwiseAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			rangeBitwiseAnd(q.m, q.n)
		}
	}
}
