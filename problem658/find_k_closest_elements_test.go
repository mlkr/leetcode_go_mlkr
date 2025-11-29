package problem658

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	arr  []int
	k, x int
	ans  []int
}{
	{
		[]int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8},
		2,
		2,
		[]int{1, 3},
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
		3,
		[]int{1, 2, 3, 4},
	},

	{
		[]int{1, 1, 2, 3, 4, 5},
		4,
		-1,
		[]int{1, 1, 2, 3},
	},
}

func Test_findClosestElements(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findClosestElements(q.arr, q.k, q.x))
	}
}

func Benchmark_findClosestElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findClosestElements(q.arr, q.k, q.x)
		}
	}
}
