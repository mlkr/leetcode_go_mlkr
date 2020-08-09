package problem434

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	s string
	ans int
}{
	{
		"Hello, my name is John",
		5,
	},

	{
		"   ",
		0,
	},
}

func Test_countSegments(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, countSegments(q.s), "输入 %s", q.s)
	}
}

func Benchmark_countSegments(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			countSegments(q.s)
		}
	}
}