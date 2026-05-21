package problem495

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	timeSeries []int
	duration   int
	ans        int
}{
	{
		[]int{1, 4},
		2,
		4,
	},

	{
		[]int{1, 2},
		2,
		3,
	},
}

func Test_findPoisonedDuration(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findPoisonedDuration(q.timeSeries, q.duration), "输入 %v", q.timeSeries)
	}
}

func Benchmark_findPoisonedDuration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findPoisonedDuration(q.timeSeries, q.duration)
		}
	}
}
