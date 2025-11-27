package problem657

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	moves string
	ans   bool
}{
	{
		"UD",
		true,
	},

	{
		"LL",
		false,
	},
}

func Test_judgeCircle(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, judgeCircle(q.moves), "输入 %v", q.moves)
	}
}

func Benchmark_judgeCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			judgeCircle(q.moves)
		}
	}
}
