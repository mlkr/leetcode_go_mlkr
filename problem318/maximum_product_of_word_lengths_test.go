package problem318

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	words []string
	ans   int
}{
	{
		[]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
		16,
	},

	{
		[]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
		4,
	},

	{
		[]string{"a", "aa", "aaa", "aaaa"},
		0,
	},

	{
		[]string{},
		0,
	},
}

func Test_maxProduct(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, maxProduct(q.words), "输入 %v", q.words)
	}
}

func Benchmark_maxProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxProduct(q.words)
		}
	}
}
