package problem664

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"aaabbb",
		2,
	},

	{
		"aba",
		2,
	},
}

func Test_strangePrinter(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, strangePrinter(q.s), "输入 %v", q.s)
	}
}

func Benchmark_strangePrinter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			strangePrinter(q.s)
		}
	}
}
