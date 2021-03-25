package problem475

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	houses, headers []int
	ans             int
}{
	{
		[]int{1, 2, 3},
		[]int{2},
		1,
	},

	{
		[]int{1, 2, 3, 4},
		[]int{1, 4},
		1,
	},

	{
		[]int{1, 5},
		[]int{2},
		3,
	},

	{
		[]int{},
		[]int{2},
		0,
	},

	{
		[]int{25921153, 510616708},
		[]int{771515668, 357571490, 44788124, 927702196, 952509530},
		153045218,
	},
}

func Test_findRadius(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findRadius(q.houses, q.headers), "输入 %v %v", q.houses, q.headers)
	}
}

func Benchmark_findRadius(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findRadius(q.houses, q.headers)
		}
	}
}
