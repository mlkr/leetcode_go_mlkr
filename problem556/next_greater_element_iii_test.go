package problem556

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		2302431,
		2303124,
	},

	{
		230241,
		230412,
	},

	{
		2147483486,
		-1,
	},

	{
		12,
		21,
	},
	{
		21,
		-1,
	},

	{
		1,
		-1,
	},
}

func Test_nextGreaterElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, nextGreaterElement(q.n), "输入 %v", q.n)
	}
}

func Benchmark_nextGreaterElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nextGreaterElement(q.n)
		}
	}
}
