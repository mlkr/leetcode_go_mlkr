package problem392

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, t string
	ans  bool
}{
	{
		"abc",
		"ahbgdc",
		true,
	},

	{
		"axc",
		"ahbgdc",
		false,
	},

	{
		"",
		"dada",
		true,
	},
}

func Test_isSubsequence(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, isSubsequence(q.s, q.t), "输入 %s", q.s)
	}
}

func Benchmark_isSubsequence(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isSubsequence(q.s, q.t)
		}
	}
}
