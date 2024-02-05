package problem547

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	isConnected [][]int
	ans         int
}{
	{
		[][]int{
			[]int{1, 1, 1},
			[]int{1, 1, 1},
			[]int{1, 1, 1},
		},
		1,
	},

	{
		[][]int{
			[]int{1, 1, 0},
			[]int{1, 1, 0},
			[]int{0, 0, 1},
		},
		2,
	},
}

func Test_findCircleNum(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findCircleNum(q.isConnected), "输入 %v", q.isConnected)
	}
}

func Test_findCircleNum2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findCircleNum2(q.isConnected), "输入 %v", q.isConnected)
	}
}

func Benchmark_findCircleNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findCircleNum(q.isConnected)
		}
	}
}

func Benchmark_findCircleNum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findCircleNum2(q.isConnected)
		}
	}
}
