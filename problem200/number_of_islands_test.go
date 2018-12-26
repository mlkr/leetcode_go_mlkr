package problem200

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para [][]byte
	ans  int
}{
	{
		[][]byte{
			[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '0'},
		},
		1,
	},

	{
		[][]byte{
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '1', '0', '0'},
			[]byte{'0', '0', '0', '1', '1'},
		},
		3,
	},
}

func Test_numIslands(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(numIslands(q.para), q.ans)
	}
}

func Benchmark_numIslands(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numIslands(q.para)
		}
	}
}
