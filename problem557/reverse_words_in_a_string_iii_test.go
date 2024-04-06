package problem557

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"Let's take LeetCode contest",
		"s'teL ekat edoCteeL tsetnoc",
	},

	{
		"Mr Ding",
		"rM gniD",
	},
}

func Test_reverseWords(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, reverseWords(q.s), "输入 %v", q.s)
	}
}

func Benchmark_reverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reverseWords(q.s)
		}
	}
}
