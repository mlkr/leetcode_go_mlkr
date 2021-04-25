package problem479

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		2,
		987,
	},

	{
		1,
		9,
	},
}

func Test_largestPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, largestPalindrome(q.n), "输入 %v", q.n)
	}
}

func Benchmark_largestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			largestPalindrome(q.n)
		}
	}
}
