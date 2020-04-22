package problem372

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	a int
	b []int
	ans int
}{
	{
		2,
		[]int{3},
		8,
	},

	{
		2,
		[]int{1,0},
		1024,
	},
}

func Test_superPow(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, superPow(q.a,q.b), "输入 %q %q", q.a, q.b)
	}
}

func Benchmark_superPow(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
				superPow(q.a,q.b)
		}
	}
}
