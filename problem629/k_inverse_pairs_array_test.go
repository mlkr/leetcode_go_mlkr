package problem629

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, k, ans int
}{
	{
		3,
		3,
		1,
	},

	{
		3,
		0,
		1,
	},

	{
		3,
		1,
		2,
	},
}

func Test_kInversePairs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, kInversePairs(q.n, q.k))
	}
}

func Benchmark_kInversePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			kInversePairs(q.n, q.k)
		}
	}
}
