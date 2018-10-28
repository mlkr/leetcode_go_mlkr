package problem139

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s        string
	wordDict []string
	ans      bool
}{
	{
		"goalspecial",
		[]string{
			"go",
			"goal",
			"goals",
			"special",
		},
		true,
	},

	{
		"aaaaaaa",
		[]string{
			"aaaa",
			"aaa",
		},
		true,
	},

	{
		"leetcode",
		[]string{
			"leet",
			"code",
		},
		true,
	},

	{
		"applepenapple",
		[]string{
			"apple",
			"pen",
		},
		true,
	},

	{
		"catsandog",
		[]string{
			"cats",
			"dog",
			"sand",
			"and",
			"cat",
		},
		false,
	},
}

func Test_wordBreak(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(wordBreak(q.s, q.wordDict), q.ans)
	}
}

func Benchmark_wordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wordBreak(q.s, q.wordDict)
		}
	}
}
