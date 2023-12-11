package problem541

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	k   int
	ans string
}{
	{
		"abcde",
		2,
		"bacde",
	},

	{
		"abcdefg",
		2,
		"bacdfeg",
	},

	{
		"abcd",
		2,
		"bacd",
	},
}

func Test_reverseStr(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, reverseStr(q.s, q.k))
	}
}

func Benchmark_reverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reverseStr(q.s, q.k)
		}
	}
}
