package problem440

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, k, ans int
}{
	{
		13,
		2,
		10,
	},

	{
		10,
		2,
		10,
	},

	{
		13,
		5,
		13,
	},
}

func Test_findKthNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findKthNumber(q.n, q.k), "%v", q.n)
	}
}

func Benchmark_findKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findKthNumber(q.n, q.k)
		}
	}
}
