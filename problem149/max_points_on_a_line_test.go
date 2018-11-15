package problem149

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []Point
	ans  int
}{
	{
		[]Point{
			Point{1, 1},
			Point{2, 2},
			Point{3, 3},
		},
		3,
	},

	{
		[]Point{
			Point{1, 1},
			Point{3, 2},
			Point{5, 3},
			Point{4, 1},
			Point{2, 3},
			Point{1, 4},
		},
		4,
	},
}

func Test_maxPoints(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(maxPoints(q.para), q.ans)
	}
}

func Benchmark_maxPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			maxPoints(q.para)
		}
	}
}
