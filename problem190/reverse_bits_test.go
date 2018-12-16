package problem190

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para uint32
	ans  uint32
}{
	{
		43261596,
		964176192,
	},

	{
		4294967293,
		3221225471,
	},
}

func Test_reverseBits(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(reverseBits(q.para), q.ans)
	}
}

func Benchmark_reverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reverseBits(q.para)
		}
	}
}
