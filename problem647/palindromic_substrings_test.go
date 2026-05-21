package problem647

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"abc",
		3,
	},

	{
		"aaa",
		6,
	},
}

func Test_countSubstrings(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, countSubstrings(q.s), "输入 %v", q.s)
	}
}

func Benchmark_countSubstrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countSubstrings(q.s)
		}
	}
}
