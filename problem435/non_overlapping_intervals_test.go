package problem435

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	intervals [][]int
	ans       int
}{
	{
		[][]int{
			{1, 2},
			{2, 3},
			{3, 4},
			{1, 3},
		},
		1,
	},

	{
		[][]int{
			{1, 2},
			{1, 2},
			{1, 2},
		},
		2,
	},

	{
		[][]int{
			{1, 2},
			{2, 3},
		},
		0,
	},
}

func Test_eraseOverlapIntervals(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, eraseOverlapIntervals(q.intervals))
	}
}

func Benchmark_eraseOverlapIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			eraseOverlapIntervals(q.intervals)
		}
	}
}
