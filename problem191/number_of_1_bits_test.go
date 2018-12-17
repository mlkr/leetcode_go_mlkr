package problem191

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para uint32
	ans  int
}{
	{
		11,
		3,
	},

	{
		128,
		1,
	},
}

func Test_hammingWeight(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(hammingWeight(q.para), q.ans)
	}
}

func Benchmark_hammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			hammingWeight(q.para)
		}
	}
}
