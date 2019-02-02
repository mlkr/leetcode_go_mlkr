package problem214

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"aacecaaa",
		"aaacecaaa",
	},

	{
		"abcd",
		"dcbabcd",
	},
}

func Test_shortestPalindrome(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(shortestPalindrome(q.s), q.ans)
	}
}

func Benchmark_shortestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			shortestPalindrome(q.s)
		}
	}
}
