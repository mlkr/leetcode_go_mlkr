package problem342

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	input int
	ans   bool
}{
	{
		2,
		false,
	},

	{
		1,
		true,
	},

	{
		0,
		false,
	},

	{
		-2147483648,
		false,
	},

	{
		16,
		true,
	},

	{
		5,
		false,
	},
}

func Test_isPowerOfFour(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, isPowerOfFour(q.input), "输入 %v", q.input)
	}
}

func Benchmark_isPowerOfFour(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPowerOfFour(q.input)
		}
	}
}
