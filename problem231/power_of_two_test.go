package problem231

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  bool
}{
	{
		1,
		true,
	},

	{
		16,
		true,
	},

	{
		218,
		false,
	},
}

func Test_isPowerOfTwo(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, isPowerOfTwo(q.para))
	}
}

func Benchmark_isPowerOfTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			isPowerOfTwo(q.para)
		}
	}
}
