package problem228

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  []string
}{
	{
		[]int{0, 1, 2, 4, 5, 7},
		[]string{"0->2", "4->5", "7"},
	},

	{
		[]int{0, 2, 3, 4, 6, 8, 9},
		[]string{"0", "2->4", "6", "8->9"},
	},

	{
		[]int{0},
		[]string{"0"},
	},

	{
		[]int{},
		nil,
	},
}

func Test_summaryRanges(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, summaryRanges(q.para))
	}
}

func Test_summaryRanges2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, summaryRanges2(q.para))
	}
}

func Benchmark_summaryRanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			summaryRanges(q.para)
		}
	}
}

func Benchmark_summaryRanges2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			summaryRanges2(q.para)
		}
	}
}
