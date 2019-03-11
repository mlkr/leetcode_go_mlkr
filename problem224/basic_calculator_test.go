package problem224

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"1 + 1",
		2,
	},

	{
		" 2-1 + 2 ",
		3,
	},

	{
		"(1+(4+5+2)-3)+(6+8)",
		23,
	},
}

func Test_calculate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, calculate(q.s))
	}
}

func Benchmark_calculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			calculate(q.s)
		}
	}
}
