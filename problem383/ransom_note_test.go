package problem383

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	r, m string
	ans  bool
}{
	{
		"a",
		"b",
		false,
	},

	{
		"aa",
		"ab",
		false,
	},

	{
		"aa",
		"aab",
		true,
	},
}

func Test_canConstruct(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, canConstruct(q.r, q.m), "输入为 %s %s", q.r, q.m)
	}
}

func Benchmark_canConstruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canConstruct(q.r, q.m)
		}
	}
}
