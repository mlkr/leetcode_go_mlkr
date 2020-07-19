package problem419

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board [][]byte
	ans   int
}{
	{
		[][]byte{},
		0,
	},

	{
		[][]byte{
			{'X', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
			{'.', '.', '.', 'X'},
		},
		2,
	},
}

func Test_countBattleships(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, countBattleships(q.board), "输入 %v", q.board)
	}
}

func Benchmark_countBattleships(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countBattleships(q.board)
		}
	}
}
