package problem481

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	n, nas int
}{
	{
		6,
		3,
	},

	{
		1,
		1,
	},
}

func Test_magicalString(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.nas, magicalString(q.n), "输入 %v", q.n)
	}
}

func Benchmark_magicalString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			magicalString(q.n)
		}

	}
}
