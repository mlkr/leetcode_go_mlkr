package problem48

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one [][]int
}

type question struct {
	para
	ans
}

func Test_rotate(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{
				[][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			ans{
				[][]int{
					{15, 13, 2, 5},
					{14, 3, 4, 1},
					{12, 6, 8, 9},
					{16, 7, 10, 11},
				},
			},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		var temp [][]int
		copy(temp, para.one)
		rotate(para.one)
		ast.Equal(ans.one, para.one, "输入 %v", temp)
	}
}
