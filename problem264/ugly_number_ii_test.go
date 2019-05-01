package problem264

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, ans int
}{
	{
		1352,
		402653184,
	},

	{
		10,
		12,
	},

	{
		6,
		6,
	},
}

func Test_nthUglyNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, nthUglyNumber(q.n), "输入%v", q.n)
	}
}

func Benchmark_nthUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			nthUglyNumber(q.n)
		}
	}
}
