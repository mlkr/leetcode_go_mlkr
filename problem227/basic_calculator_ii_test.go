package problem227

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para string
	ans  int
}{
	{
		"3+2*2",
		7,
	},

	{
		" 3/2 ",
		1,
	},

	{
		" 3+5 / 2 ",
		5,
	},

	{
		" 3-5 / 2 ",
		1,
	},
}

func Test_calculate(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, calculate(q.para))
	}
}

func Benchmark_calculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			calculate(q.para)
		}
	}
}
