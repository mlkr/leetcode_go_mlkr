package problem509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		2,
		1,
	},

	{
		3,
		2,
	},

	{
		4,
		3,
	},
}

func Test_fib(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, fib(q.n), "输入 %v", q.n)
	}
}

func Benchmark_fib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			fib(q.n)
		}
	}
}
