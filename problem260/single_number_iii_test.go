package problem260

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para, ans []int
}{
	{
		[]int{1, 2, 1, 3, 2, 5},
		[]int{3, 5},
	},

	{
		[]int{3, 2, 1, 3, 2, 5},
		[]int{1, 5},
	},
}

func Test_singleNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, singleNumber(q.para))
	}
}

func Test_singleNumber2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, singleNumber2(q.para))
	}
}

func Test_singleNumber3(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, singleNumber3(q.para))
	}
}

func Benchmark_singleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber(q.para)
		}
	}
}

func Benchmark_singleNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber2(q.para)
		}
	}
}

func Benchmark_singleNumber3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			singleNumber3(q.para)
		}
	}
}
