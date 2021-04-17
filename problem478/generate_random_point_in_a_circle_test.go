package problem478

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	r, x, y float64
}{
	{
		1.,
		0.,
		0.,
	},

	{
		3332138.342,
		321.676,
		87976.4,
	},
}

func Test_Solution(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		so := Constructor(q.r, q.x, q.y)
		for i := 0; i < 1000; i++ {
			res := so.RandPoint()
			x, y := res[0], res[1]
			ast.True((x-so.x)*(x-so.x)+(y-so.y)*(y-so.y) <= so.r*so.r)
		}
	}
}

func Benchmark_Solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			so := Constructor(q.r, q.x, q.y)
			for i := 0; i < 1000; i++ {
				so.RandPoint()
			}
		}
	}
}
