package problem352

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	ans  [][][]int
}{
	{
		[]int{1, 3, 7, 2, 6, -2, -2, -3, 8, 13, 10, 11},
		[][][]int{
			[][]int{[]int{1, 1}},
			[][]int{[]int{1, 1}, []int{3, 3}},
			[][]int{[]int{1, 1}, []int{3, 3}, []int{7, 7}},
			[][]int{[]int{1, 3}, []int{7, 7}},
			[][]int{[]int{1, 3}, []int{6, 7}},
			[][]int{[]int{-2, -2}, []int{1, 3}, []int{6, 7}},
			[][]int{[]int{-2, -2}, []int{1, 3}, []int{6, 7}},
			[][]int{[]int{-3, -2}, []int{1, 3}, []int{6, 7}},
			[][]int{[]int{-3, -2}, []int{1, 3}, []int{6, 8}},
			[][]int{[]int{-3, -2}, []int{1, 3}, []int{6, 8}, []int{13, 13}},
			[][]int{[]int{-3, -2}, []int{1, 3}, []int{6, 8}, []int{10, 10}, []int{13, 13}},
			[][]int{[]int{-3, -2}, []int{1, 3}, []int{6, 8}, []int{10, 11}, []int{13, 13}},
		},
	},
}

func Test_SummaryRanges(t *testing.T) {
	ast := assert.New(t)

	summaryRanges := Constructor()
	for _, q := range questions {
		for k, n := range q.nums {
			summaryRanges.AddNum(n)
			ast.ElementsMatch(summaryRanges.GetIntervals(), q.ans[k], "输入 %d", n)
		}
	}
}

func Benchmark_summaryRanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		summaryRanges := Constructor()
		for _, q := range questions {
			for _, n := range q.nums {
				summaryRanges.AddNum(n)
				summaryRanges.GetIntervals()
			}
		}
	}
}
