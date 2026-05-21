package problem593

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	p1, p2, p3, p4 []int
	ans            bool
}{
	{
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		[]int{0, 0},
		false,
	},

	{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
		true,
	},

	{
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
		[]int{0, 1},
		true,
	},

	{
		[]int{0, 0},
		[]int{1, 1},
		[]int{1, 0},
		[]int{0, 12},
		false,
	},

	{
		[]int{1, 0},
		[]int{-1, 0},
		[]int{0, 1},
		[]int{0, -1},
		true,
	},
}

func Test_validSquare(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, validSquare(q.p1, q.p2, q.p3, q.p4))
	}
}

func Benchmark_validSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			validSquare(q.p1, q.p2, q.p3, q.p4)
		}
	}
}
