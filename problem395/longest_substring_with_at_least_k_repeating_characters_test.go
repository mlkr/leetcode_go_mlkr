package problem395

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	k   int
	ans int
}{
	{
		"bbaaacbd",
		3,
		3,
	},

	{
		"aaabb",
		3,
		3,
	},

	{
		"ababbc",
		2,
		5,
	},
}

func Test_longestSubstring(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, longestSubstring(q.s, q.k), "输入 %s", q.s)
	}
}

func Benchmark_longestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			longestSubstring(q.s, q.k)
		}
	}
}
