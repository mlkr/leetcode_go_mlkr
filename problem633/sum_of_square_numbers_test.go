package problem633

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	c   int
	ans bool
}{
	{
		2,
		true,
	},

	{
		5,
		true,
	},

	{
		3,
		false,
	},
}

func Test_judgeSquareSum(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, judgeSquareSum(q.c), "输入 %v", q.c)
	}
}

func Benchmark_judgeSquareSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			judgeSquareSum(q.c)
		}
	}
}
