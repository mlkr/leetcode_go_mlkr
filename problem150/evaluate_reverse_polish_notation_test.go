package problem150

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	tokens []string
	ans    int
}{
	{
		[]string{"2", "1", "+", "3", "*"},
		9,
	},

	{
		[]string{"4", "13", "5", "/", "+"},
		6,
	},

	{
		[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
		22,
	},
}

func Test_evalRPN(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(evalRPN(q.tokens), q.ans)
	}
}

func Benchmark_evalRPN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			evalRPN(q.tokens)
		}
	}
}
