package problem316

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	str, ans string
}{
	{
		"bcabc",
		"abc",
	},

	{
		"cbacdcbc",
		"acdb",
	},
}

func Test_removeDuplicateLetters(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, removeDuplicateLetters(q.str), "输入 %v", q.str)
	}
}

func Test_removeDuplicateLetters2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, removeDuplicateLetters2(q.str), "输入 %v", q.str)
	}
}

func Benchmark_removeDuplicateLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			removeDuplicateLetters(q.str)
		}
	}
}

func Benchmark_removeDuplicateLetters2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			removeDuplicateLetters2(q.str)
		}
	}
}
