package problem583

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	word1, word2 string
	ans          int
}{
	{
		"sea",
		"eat",
		2,
	},

	{
		"leetcode",
		"etco",
		4,
	},
}

func Test_minDistance(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minDistance(q.word1, q.word2), "%v , %v", q.word1, q.word2)
	}
}

func Benchmark_minDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			minDistance(q.word1, q.word2)
		}
	}
}
