package problem521

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	a, b string
	ans  int
}{
	{
		"aba",
		"cdc",
		3,
	},

	{
		"aaa",
		"bbb",
		3,
	},

	{
		"aaab",
		"bbb",
		4,
	},

	{
		"aaa",
		"bbba",
		4,
	},

	{
		"aaa",
		"aaa",
		-1,
	},
}

func Test_findLUSlength(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findLUSlength(q.a, q.b), "输入 %v %v", q.a, q.b)
	}
}

func Benchmark_findLUSlength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLUSlength(q.a, q.b)
		}
	}
}
