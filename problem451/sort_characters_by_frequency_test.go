package problem451

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	s, ans string
}{
	{
		"tree",
		"eert",
	},

	{
		"cccaaa",
		"aaaccc",
	},

	{
		"Aabb",
		"bbAa",
	},
}

func Test_frequencySort(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, frequencySort(q.s), "输入 %v", q.ans)
	}
}

func Test_frequencySort2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, frequencySort2(q.s), "输入 %v", q.ans)
	}
}

func Benchmark_frequencySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			frequencySort(q.s)
		}
	}
}

func Benchmark_frequencySort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			frequencySort2(q.s)
		}
	}
}
