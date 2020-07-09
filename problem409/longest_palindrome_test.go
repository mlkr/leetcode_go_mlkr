package problem409

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"abccccdd",
		7,
	},

	{
		"",
		0,
	},
}

func Test_longestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, longestPalindrome(q.s), "输入 %s", q.s)
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			longestPalindrome(q.s)
		}
	}
}
