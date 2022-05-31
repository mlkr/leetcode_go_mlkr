package problem483

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans string
}{

	{
		"9",
		"8",
	},

	{
		"11",
		"10",
	},

	{
		"13",
		"3",
	},

	{
		"4681",
		"8",
	},
}

func Test_smallestGoodBase(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, smallestGoodBase(q.n), "输入 %v", q.n)
	}
}

func Test_isFound(t *testing.T) {
	ast := assert.New(t)

	k := 3
	m := 2
	num := 13
	ast.True(isFound(k, m, num), "输入 %v, %v, %v", k, m, num)

}

func Benchmark_smallestGoodBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			smallestGoodBase(q.n)
		}
	}
}
