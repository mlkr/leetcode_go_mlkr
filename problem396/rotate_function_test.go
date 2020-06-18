package problem396

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	A   []int
	ans int
}{
	{
		[]int{4, 3, 2, 6},
		26,
	},
}

func Test_maxRotateFunction(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, maxRotateFunction(q.A))
	}
}

func Benchmark_maxRotateFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxRotateFunction(q.A)
		}
	}
}
