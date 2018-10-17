package problem132

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para string
	ans  int
}{
	{
		"ababababababababababababcbabababababababababababa",
		0,
	},

	{
		"aaabaa",
		1,
	},

	{
		"aab",
		1,
	},
}

func Test_minCut(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(minCut(q.para), q.ans, "输入 %v", q.para)
	}
}

func Benchmark_minCut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minCut(q.para)
		}
	}
}
