package problem263

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num int
	ans bool
}{
	{
		0,
		false,
	},

	{
		5,
		true,
	},

	{
		6,
		true,
	},

	{
		8,
		true,
	},

	{
		14,
		false,
	},
}

func Test_isUgly(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isUgly(q.num))
	}
}

func Benchmark_isUgly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isUgly(q.num)
		}
	}
}
