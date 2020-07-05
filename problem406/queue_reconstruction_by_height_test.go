package problem406

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	people [][]int
	ans    [][]int
}{
	{
		[][]int{
			[]int{7, 0},
			[]int{4, 4},
			[]int{7, 1},
			[]int{5, 0},
			[]int{6, 1},
			[]int{5, 2},
		},

		[][]int{
			[]int{5, 0},
			[]int{7, 0},
			[]int{5, 2},
			[]int{6, 1},
			[]int{4, 4},
			[]int{7, 1},
		},
	},
}

func Test_reconstructQueue(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, reconstructQueue(q.people), "输入 %v", q.people)
	}
}

func Test_reconstructQueue2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, reconstructQueue2(q.people), "输入 %v", q.people)
	}
}

func Benchmark_reconstructQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reconstructQueue(q.people)
		}
	}
}

func Benchmark_reconstructQueue2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			reconstructQueue2(q.people)
		}
	}
}
