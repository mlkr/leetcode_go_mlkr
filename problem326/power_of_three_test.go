package problem326

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	num int
	ans bool
}{
	{
		27,
		true,
	},

	{
		0,
		false,
	},

	{
		9,
		true,
	},

	{
		45,
		false,
	},
}

func Test_isPowerOfThree(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, isPowerOfThree(q.num))
	}
}

func Benchmark_isPowerOfThree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPowerOfThree(q.num)
		}
	}
}
