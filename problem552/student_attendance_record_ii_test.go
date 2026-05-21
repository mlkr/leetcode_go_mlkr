package problem552

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		4,
		43,
	},

	{
		3,
		19,
	},

	{
		2,
		8,
	},

	{
		1,
		3,
	},

	{
		10101,
		183236316,
	},
}

func Test_checkRecord(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, checkRecord(q.n), "输入 %v", q.n)
	}
}

func Benchmark_checkRecord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			checkRecord(q.n)
		}
	}
}
