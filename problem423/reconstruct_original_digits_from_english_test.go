package problem423

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"nnei",
		"9",
	},

	{
		"owoztneoer",
		"012",
	},

	{
		"fviefuro",
		"45",
	},

	{
		"sixeightseven",
		"678",
	},
}

// zero one two three four five six seven eight nine

func Test_originalDigits(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, originalDigits(q.s), "输入 %v", q.s)
	}
}

func Benchmark_originalDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			originalDigits(q.s)
		}
	}
}
