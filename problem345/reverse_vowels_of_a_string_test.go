package problem345

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"aA",
		"Aa",
	},

	{
		"hello",
		"holle",
	},

	{
		"leetcode",
		"leotcede",
	},
}

func Test_reverseVowels(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, reverseVowels(q.s), "输入 %v", q.s)
	}
}

func Benchmark_reverseVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reverseVowels(q.s)
		}
	}
}
