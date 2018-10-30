package problem140

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s        string
	wordDict []string
	ans      []string
}{
	{
		"catsanddog",
		[]string{
			"cat",
			"cats",
			"and",
			"sand",
			"dog",
		},
		[]string{
			"cat sand dog",
			"cats and dog",
		},
	},

	{
		"pineapplepenapple",
		[]string{
			"apple",
			"pen",
			"applepen",
			"pine",
			"pineapple",
		},
		[]string{
			"pine apple pen apple",
			"pine applepen apple",
			"pineapple pen apple",
		},
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
		nil,
	},
}

func Test_wordBreak(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(wordBreak(q.s, q.wordDict), q.ans)
	}
}

func Test_wordBreak2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(wordBreak2(q.s, q.wordDict), q.ans)
	}
}

func Test_wordBreak3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(wordBreak3(q.s, q.wordDict), q.ans)
	}
}

func Benchmark_wordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wordBreak(q.s, q.wordDict)
		}
	}
}

func Benchmark_wordBreak2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wordBreak2(q.s, q.wordDict)
		}
	}
}

func Benchmark_wordBreak3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			wordBreak3(q.s, q.wordDict)
		}
	}
}
