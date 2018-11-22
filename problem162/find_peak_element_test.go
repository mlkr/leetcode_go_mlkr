package problem162

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{1},
		0,
	},

	{
		[]int{1, 2, 3, 1},
		2,
	},

	{
		[]int{1, 2, 1, 3, 5, 6, 4},
		1,
	},
}

func Test_findPeakElement(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(findPeakElement(q.para), q.ans)
	}
}

func Benchmark_findPeakElement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findPeakElement(q.para)
		}

	}
}
