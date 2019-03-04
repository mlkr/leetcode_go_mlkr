package problem221

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	matrix [][]byte
	ans    int
}{
	{
		[][]byte{
			[]byte{'1', '0', '1', '0', '0'},
			[]byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'},
			[]byte{'1', '0', '0', '1', '0'},
		},
		4,
	},
}

func Test_maximalSquare(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maximalSquare(q.matrix))
	}
}

func Benchmark_maximalSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maximalSquare(q.matrix)
		}

	}
}
