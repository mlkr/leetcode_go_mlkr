package problem668

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	m, n, k, ans int
}{
	{
		16, 50, 378,
		152,
	},

	{
		3, 3, 5,
		3,
	},

	{
		2, 3, 6,
		6,
	},
}

func Test_findKthNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findKthNumber(q.m, q.n, q.k))
	}
}

func Benchmark_findKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findKthNumber(q.m, q.n, q.k)
		}
	}
}
