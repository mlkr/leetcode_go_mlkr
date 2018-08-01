package problem57

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []Interval
	two Interval
}

type ans struct {
	one []Interval
}

type question struct {
	para
	ans
}

func TestInsert(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]Interval{
				Interval{1, 5},
			},
				Interval{0, 6},
			},
			ans{[]Interval{
				Interval{0, 6},
			}},
		},

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{6, 9},
			},
				Interval{2, 5},
			},
			ans{[]Interval{
				Interval{1, 5},
				Interval{6, 9},
			}},
		},

		question{
			para{[]Interval{
				Interval{1, 2},
				Interval{3, 5},
				Interval{6, 7},
				Interval{8, 10},
				Interval{12, 16},
			},
				Interval{4, 8},
			},
			ans{[]Interval{
				Interval{1, 2},
				Interval{3, 10},
				Interval{12, 16},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, insert(para.one, para.two), "输入 %v", para)
	}
}
