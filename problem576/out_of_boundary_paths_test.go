package problem576

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	m, n, maxMove, startRow, startCol int
	ans                               int
}{
	{
		2, 2, 2, 0, 0,
		6,
	},

	{
		1, 3, 3, 0, 1,
		12,
	},
}

func Test_findPaths(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findPaths(q.m, q.n, q.maxMove, q.startRow, q.startCol))
	}
}

func Benchmark_findPaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findPaths(q.m, q.n, q.maxMove, q.startRow, q.startCol)
		}
	}
}
