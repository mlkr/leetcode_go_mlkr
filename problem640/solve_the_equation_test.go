package problem640

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	equation, ans string
}{
	{
		"-x=-1",
		"x=1",
	},

	{
		"2x+3x-6x=x+2",
		"x=-1",
	},

	{
		"x+5-3+x=6+x-2",
		"x=2",
	},

	{
		"x=x",
		"Infinite solutions",
	},

	{
		"2x=x",
		"x=0",
	},
}

func Test_solveEquation(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, solveEquation(q.equation), "输入 %v", q.equation)
	}
}

func Benchmark_solveEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			solveEquation(q.equation)
		}
	}
}
