package problem357

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n   int
	ans int
}{
	{
		0,
		1,
	},

	{
		1,
		10,
	},

	{
		2,
		91,
	},

	{
		11,
		8877691,
	},

	{
		12,
		8877691,
	},
}

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, countNumbersWithUniqueDigits(q.n), "输入 %q", q.n)
	}
}

func Benchmark_countNumbersWithUniqueDigits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			countNumbersWithUniqueDigits(q.n)
		}
	}
}
