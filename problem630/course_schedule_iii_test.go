package problem630

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	courses [][]int
	ans     int
}{
	{
		[][]int{
			[]int{100, 200},
			[]int{200, 1300},
			[]int{1000, 1250},
			[]int{2000, 3200},
		},
		3,
	},

	{
		[][]int{
			[]int{1, 2},
		},
		1,
	},

	{
		[][]int{
			[]int{3, 2},
			[]int{4, 3},
		},
		0,
	},
}

func Test_scheduleCourse(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, scheduleCourse(q.courses), "输入 %v", q.courses)
	}
}

func Benchmark_scheduleCourse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			scheduleCourse(q.courses)
		}
	}
}
