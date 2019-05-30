package problem292

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n   int
	ans bool
}{
	{
		4,
		false,
	},
}

func Test_canWinNim(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, canWinNim(q.n))
	}
}

func Benchmark_canWinNim(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			canWinNim(q.n)
		}
	}
}
