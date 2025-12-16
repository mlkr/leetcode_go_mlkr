package problem661

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	img, ans [][]int
}{
	{
		[][]int{
			[]int{1, 1, 1},
			[]int{1, 0, 1},
			[]int{1, 1, 1},
		},
		[][]int{
			[]int{0, 0, 0},
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
	},

	{
		[][]int{
			[]int{100, 200, 100},
			[]int{200, 50, 200},
			[]int{100, 200, 100},
		},
		[][]int{
			[]int{137, 141, 137},
			[]int{141, 138, 141},
			[]int{137, 141, 137},
		},
	},
}

func Test_imageSmoother(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, imageSmoother(q.img), "输入 %v", q.img)
	}
}

func Benchmark_imageSmoother(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			imageSmoother(q.img)
		}
	}
}
