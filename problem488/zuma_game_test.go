package problem488

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board, hand string
	ans         int
}{

	{
		"RRWWRRBBRR",
		"WB",
		2,
	},

	{
		"WRRBBW",
		"RB",
		-1,
	},

	{
		"WWRRBBWW",
		"WRBRW",
		2,
	},

	{
		"G",
		"GGGGG",
		2,
	},
}

func Test_findMinStep(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMinStep(q.board, q.hand), "输入 %v", q.board)
	}
}

func Benchmark_findMinStep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMinStep(q.board, q.hand)
		}
	}
}
