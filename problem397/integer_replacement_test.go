package problem397

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		1234,
		14,
	},

	{
		8,
		3,
	},

	{
		7,
		4,
	},

	{
		65535,
		17,
	},

	{
		65536,
		16,
	},
}

func Test_integerReplacement(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, integerReplacement(q.n), "输入 %d", q.n)
	}
}

func Benchmark_integerReplacement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			integerReplacement(q.n)
		}
	}
}
