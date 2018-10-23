package problem135

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  int
}{
	{
		[]int{1, 0, 2},
		5,
	},

	{
		[]int{1, 2, 2},
		4,
	},
}

func Test_candy(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(candy(q.para), q.ans, "输入 %v", q.para)
	}
}

func Benchmark_candy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			candy(q.para)
		}
	}
}
