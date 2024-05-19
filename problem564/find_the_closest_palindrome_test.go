package problem564

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans string
}{
	// {
	// 	"123",
	// 	"121",
	// },

	// {
	// 	"1",
	// 	"0",
	// },

	{
		"6",
		"5",
	},
}

func Test_nearestPalindromic(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, nearestPalindromic(q.n), "输入 %v", q.n)
	}
}

func Benchmark_nearestPalindromic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nearestPalindromic(q.n)
		}
	}
}
