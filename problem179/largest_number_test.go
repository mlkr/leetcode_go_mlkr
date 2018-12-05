package problem179

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para []int
	ans  string
}{
	{
		[]int{0, 0},
		"0",
	},

	{
		[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		"9876543210",
	},

	{
		[]int{12, 121},
		"12121",
	},

	{
		[]int{121, 12},
		"12121",
	},

	{
		[]int{10, 2},
		"210",
	},

	{
		[]int{3, 30, 34, 5, 9},
		"9534330",
	},
}

func Test_largestNumber(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(largestNumber(q.para), q.ans)
	}
}

func Test_largestNumber2(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(largestNumber2(q.para), q.ans)
	}
}

func Test_compare(t *testing.T) {
	ast := assert.New(t)
	ast.Equal(compare("34", "3"), 1, "343")
	ast.Equal(compare("121", "12"), -1, "12112")
}

func Benchmark_largestNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			largestNumber(q.para)
		}
	}
}

func Benchmark_largestNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			largestNumber2(q.para)
		}
	}
}
