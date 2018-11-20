package problem154

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{1, 3, 5},
		1,
	},

	{
		[]int{2, 2, 2, 0, 1},
		0,
	},

	{
		[]int{1},
		1,
	},

	{
		[]int{3, 4, 5, 1, 2},
		1,
	},

	{
		[]int{4, 5, 6, 7, 0, 1, 2},
		0,
	},
}

func Test_findMin(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(findMin(q.para), q.ans)
	}
}

func Benchmark_findMin(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, q := range questions {
			findMin(q.para)
		}
	}
}
