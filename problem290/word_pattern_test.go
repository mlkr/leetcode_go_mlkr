package problem290

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	pattern, str string
	ans          bool
}{
	{
		"abba",
		"dog cat cat dog",
		true,
	},

	{
		"abba",
		"dog cat cat fish",
		false,
	},

	{
		"aaaa",
		"dog cat cat dog",
		false,
	},

	{
		"abba",
		"dog dog dog dog",
		false,
	},

	{
		"abb",
		"dog dog dog dog",
		false,
	},
}

func Test_wordPattern(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, wordPattern(q.pattern, q.str), "输入 %s", q.str)
	}
}

func Benchmark_wordPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wordPattern(q.pattern, q.str)
		}
	}
}
