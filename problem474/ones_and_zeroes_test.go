package problem474

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	strs []string
	m, n int
	ans  int
}{

	{
		[]string{"10", "0001", "111001", "1", "0"},
		5,
		3,
		4,
	},

	{
		[]string{"111", "1000", "1000", "1000"},
		9,
		3,
		3,
	},

	{
		[]string{"10", "0", "1"},
		1,
		1,
		2,
	},

	{
		[]string{"1", "0", "10", "0001", "111001"},
		5,
		3,
		4,
	},
}

func Test_findMaxForm(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, findMaxForm(q.strs, q.m, q.n), "输入 %v %v %v", q.strs, q.m, q.n)
	}
}

func Benchmark_findMaxForm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findMaxForm(q.strs, q.m, q.n)
		}
	}
}
