package problem639

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s   string
	ans int
}{
	{
		"11106",
		2,
	},

	{
		"*",
		9,
	},

	{
		"1*",
		18,
	},

	{
		"2*",
		15,
	},

	{
		"*********",
		291868912,
	},
}

func Test_numDecodings(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, numDecodings(q.s), "输入 %v", q.s)
	}
}

func Benchmark_numDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			numDecodings(q.s)
		}
	}
}
