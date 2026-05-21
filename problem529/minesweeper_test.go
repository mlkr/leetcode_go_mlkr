package problem529

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board [][]byte
	click []int
	ans   [][]byte
}{
	{
		[][]byte{
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'M', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
			{'E', 'E', 'E', 'E', 'E'},
		},
		[]int{3, 0},
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	},

	{
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'M', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
		[]int{1, 2},
		[][]byte{
			{'B', '1', 'E', '1', 'B'},
			{'B', '1', 'X', '1', 'B'},
			{'B', '1', '1', '1', 'B'},
			{'B', 'B', 'B', 'B', 'B'},
		},
	},
}

func Test_updateBoard(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, updateBoard(q.board, q.click))
	}
}

func Benchmark_updateBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			updateBoard(q.board, q.click)
		}
	}
}
