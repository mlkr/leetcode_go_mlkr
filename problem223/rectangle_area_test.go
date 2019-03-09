package problem223

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	a, b, c, d, e, f, g, h int
	ans                    int
}{
	{
		-3, 0, 3, 4, 0, -1, 9, 2,
		45,
	},
}

func Test_computeArea(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, computeArea(q.a, q.b, q.c, q.d, q.e, q.f, q.g, q.h))
	}
}

func Benchmark_computeArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			computeArea(q.a, q.b, q.c, q.d, q.e, q.f, q.g, q.h)
		}
	}
}
