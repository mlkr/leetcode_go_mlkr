package problem472

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	words, ans []string
}{
	{
		[]string{
			"a", "b", "ab", "abc",
		},
		[]string{
			"ab",
		},
	},

	{
		[]string{
			"a", "b", "ab", "abc", "aab",
		},
		[]string{
			"ab", "aab",
		},
	},

	{
		[]string{
			"cat", "cats", "catsdogcats", "dog", "dogcatsdog",
			"hippopotamuses", "rat", "ratcatdogcat",
		},
		[]string{
			"catsdogcats", "dogcatsdog", "ratcatdogcat",
		},
	},

	{
		[]string{
			"cat", "dog", "catdog",
		},
		[]string{
			"catdog",
		},
	},
}

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, findAllConcatenatedWordsInADict(q.words), "输入 %v", q.words)
	}
}

func Benchmark_findAllConcatenatedWordsInADict(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findAllConcatenatedWordsInADict(q.words)
		}
	}
}
