package problem461

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	x, y, ans int
}{
	{
		1,
		4,
		2,
	},
}

func Test_hammingDistance(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, hammingDistance(q.x, q.y), "x %d, y %d", q.x, q.y)
	}
}

func Benchmark_hammingDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			hammingDistance(q.x, q.y)
		}
	}
}
