package problem130

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board [][]byte
	ans   [][]byte
}{
	{
		[][]byte{
			[]byte{'O', 'X', 'O', 'O', 'X', 'X'},
			[]byte{'O', 'X', 'X', 'X', 'O', 'X'},
			[]byte{'X', 'O', 'O', 'X', 'O', 'O'},
			[]byte{'X', 'O', 'X', 'X', 'X', 'X'},
			[]byte{'O', 'O', 'X', 'O', 'X', 'X'},
			[]byte{'X', 'X', 'O', 'O', 'O', 'O'},
		},
		[][]byte{
			[]byte{'O', 'X', 'O', 'O', 'X', 'X'},
			[]byte{'O', 'X', 'X', 'X', 'O', 'X'},
			[]byte{'X', 'O', 'O', 'X', 'O', 'O'},
			[]byte{'X', 'O', 'X', 'X', 'X', 'X'},
			[]byte{'O', 'O', 'X', 'O', 'X', 'X'},
			[]byte{'X', 'X', 'O', 'O', 'O', 'O'},
		},
	},

	{
		[][]byte{},
		[][]byte{},
	},

	{
		[][]byte{
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'O', 'O', 'X'},
			[]byte{'X', 'X', 'O', 'X'},
			[]byte{'X', 'O', 'X', 'X'},
		},

		[][]byte{
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'X', 'X', 'X'},
			[]byte{'X', 'O', 'X', 'X'},
		},
	},
}

func TestSolve(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		solve(q.board)
		ast.Equal(q.board, q.ans)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			solve(q.board)
		}
	}
}
