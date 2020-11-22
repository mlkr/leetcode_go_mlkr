package problem459

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans bool
}{
	// {
	// 	"abab",
	// 	true,
	// },

	{
		"aba",
		false,
	},

	{
		"abacd",
		false,
	},

	{
		"abcabcabcabc",
		true,
	},

	{
		"a",
		false,
	},
}

func Test_repeatedSubstringPattern(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, repeatedSubstringPattern(q.s), "输入 %v", q.s)
	}
}

func Test_repeatedSubstringPattern2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, repeatedSubstringPattern2(q.s), "输入 %v", q.s)
	}
}

func Benchmark_repeatedSubstringPattern(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			repeatedSubstringPattern(q.s)
		}
	}
}

func Benchmark_repeatedSubstringPattern2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			repeatedSubstringPattern2(q.s)
		}
	}
}
