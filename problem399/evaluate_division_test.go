package problem399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	equations [][]string
	values    []float64
	queries   [][]string
	ans       []float64
}{
	{
		[][]string{
			[]string{"a", "b"},
			[]string{"b", "c"},
		},
		[]float64{
			2.0,
			3.0,
		},
		[][]string{
			[]string{"a", "c"},
			[]string{"b", "a"},
			[]string{"a", "e"},
			[]string{"a", "a"},
			[]string{"x", "x"},
		},
		[]float64{
			6.0,
			0.5,
			-1.0,
			1.0,
			-1.0,
		},
	},
}

func Test_calcEquation(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, calcEquation(q.equations, q.values, q.queries), "输入 %s", q.equations)
	}
}

func Benchmark_calcEquation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			calcEquation(q.equations, q.values, q.queries)
		}
	}
}
