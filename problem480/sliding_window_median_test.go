package problem480

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	nums []int
	k    int
	ans  []float64
}{
	{
		[]int{1},
		1,
		[]float64{1},
	},

	{
		[]int{1, 4, 2, 3},
		4,
		[]float64{2.5},
	},

	{
		[]int{1, 3, -1, -3, 5, 3, 6, 7},
		3,
		[]float64{1, -1, -1, 3, 5, 6},
	},

	{
		[]int{1, 2, 3, 4, 2, 3, 1, 4, 2},
		3,
		[]float64{2, 3, 3, 3, 2, 3, 2},
	},
}

func Test_medianSlidingWindow(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, medianSlidingWindow(q.nums, q.k), "输入 %v %v", q.nums, q.k)
	}
}

func Benchmark_medianSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			medianSlidingWindow(q.nums, q.k)
		}
	}
}

func Test_pop(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	k := 14
	w := getWindow(nums, k)
	w.minPQ.Pop()
	w.maxPQ.Pop()
}
