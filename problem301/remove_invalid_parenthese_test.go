package problem301

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans []string
}{
	{
		"()())()",
		[]string{
			"()()()",
			"(())()",
		},
	},

	{
		"(a)())()",
		[]string{
			"(a)()()",
			"(a())()",
		},
	},

	{
		")(",
		[]string{
			"",
		},
	},
}

func Test_removeInvalidParentheses(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, removeInvalidParentheses(q.s))
	}
}

func Benchmark_removeInvalidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			removeInvalidParentheses(q.s)
		}
	}
}
