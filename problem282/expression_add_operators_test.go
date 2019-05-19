package problem282

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num    string
	target int
	ans    []string
}{
	{
		"123",
		6,
		[]string{
			"1+2+3",
			"1*2*3",
		},
	},

	{
		"232",
		8,
		[]string{
			"2*3+2",
			"2+3*2",
		},
	},

	{
		"105",
		5,
		[]string{
			"1*0+5",
			"10-5",
		},
	},

	{
		"00",
		0,
		[]string{
			"0+0",
			"0-0",
			"0*0",
		},
	},

	{
		"3456237490",
		9191,
		[]string{},
	},
}

func Test_addOperators(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, addOperators(q.num, q.target))
	}
}

func Benchmark_addOperators(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			addOperators(q.num, q.target)
		}
	}
}
