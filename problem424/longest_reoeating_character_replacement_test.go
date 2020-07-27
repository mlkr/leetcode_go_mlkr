package problem424

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	k   int
	ans int
}{
	{
		"ABAB",
		2,
		4,
	},

	{
		"AABABBA",
		1,
		4,
	},
}

func Test_characterReplacement(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, characterReplacement(q.s, q.k), "输入%v", q.s)
	}
}

func Benchmark_characterReplacement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			characterReplacement(q.s, q.k)
		}
	}
}
