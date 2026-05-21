package problem520

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	word string
	ans  bool
}{
	{
		"USA",
		true,
	},

	{
		"FlaG",
		false,
	},

	{
		"flaG",
		false,
	},

	{
		"leetcode",
		true,
	},

	{
		"Google",
		true,
	},
}

func Test_detectCapitalUse(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, detectCapitalUse(q.word), "输入 %v", q.word)
	}
}

func Benchmark_detectCapitalUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			detectCapitalUse(q.word)
		}
	}
}
