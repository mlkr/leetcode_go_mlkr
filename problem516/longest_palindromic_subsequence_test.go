package problem516

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"bbbab",
		4,
	},

	{
		"cbbd",
		2,
	},
}

func Test_longestPalindromeSubseq(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, longestPalindromeSubseq(q.s), "输入 %v", q.s)
	}
}

func Benchmark_longestPalindromeSubseq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			longestPalindromeSubseq(q.s)
		}
	}
}
