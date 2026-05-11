package problem672

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, presses int
	ans        int
}{
	{
		3, 2,
		7,
	},

	{
		1, 1,
		2,
	},

	{
		2, 1,
		3,
	},

	{
		3, 1,
		4,
	},
}

func Test_flipLights(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, flipLights(q.n, q.presses))
	}
}

func Benchmark_flipLights(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			flipLights(q.n, q.presses)
		}
	}
}
