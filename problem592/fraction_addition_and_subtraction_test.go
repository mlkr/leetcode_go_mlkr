package problem592

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	expression, ans string
}{
	{
		"-1/2+1/2",
		"0/1",
	},

	{
		"-1/2+1/2+1/3",
		"1/3",
	},

	{
		"1/3-1/2",
		"-1/6",
	},
}

func Test_fractionAddition(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, fractionAddition(q.expression), "输入 %v", q.expression)
	}
}

func Benchmark_fractionAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			fractionAddition(q.expression)
		}
	}
}
