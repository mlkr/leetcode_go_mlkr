package problem172

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para int
	ans  int
}{
	{
		10,
		2,
	},

	{
		30,
		7,
	},

	{
		3,
		0,
	},

	{
		5,
		1,
	},
}

func Test_trailingZeroes(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(trailingZeroes(q.para), q.ans)
	}
}

func Benchmark_trailingZeroes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			trailingZeroes(q.para)
		}
	}
}
