package problem394

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"3[a]2[bc]",
		"aaabcbc",
	},

	{
		"3[a2[c]]",
		"accaccacc",
	},

	{
		"2[abc]3[cd]ef",
		"abcabccdcdcdef",
	},

	{
		"abc3[cd]xyz",
		"abccdcdcdxyz",
	},

	{
		"10[a]",
		"aaaaaaaaaa",
	},
}

func Test_decodeString(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, decodeString(q.s), "输入 %s", q.s)
	}
}

func Benchmark_decodeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			decodeString(q.s)
		}
	}
}
