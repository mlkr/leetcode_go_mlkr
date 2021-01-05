package problem466

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s1  string
	n1  int
	s2  string
	n2  int
	ans int
}{
	{
		"acb",
		1,
		"acb",
		1,
		1,
	},

	{
		"acb",
		4,
		"ab",
		2,
		2,
	},

	{
		"a",
		1,
		"ab",
		2,
		0,
	},
}

func Test_getMaxRepetitions(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, getMaxRepetitions(q.s1, q.n1, q.s2, q.n2), "输入 %v", q.s1)
	}
}

func Test_getMaxRepetitions2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, getMaxRepetitions2(q.s1, q.n1, q.s2, q.n2), "输入 %v", q.s1)
	}
}

func Benchmark_getMaxRepetitions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getMaxRepetitions(q.s1, q.n1, q.s2, q.n2)
		}
	}
}

func Benchmark_getMaxRepetitions2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			getMaxRepetitions2(q.s1, q.n1, q.s2, q.n2)
		}
	}
}
