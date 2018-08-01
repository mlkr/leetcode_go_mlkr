package problem56

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []Interval
}

type ans struct {
	one []Interval
}

type question struct {
	para
	ans
}

func TestMerge(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{[]Interval{
				Interval{1, 3},
				Interval{2, 6},
				Interval{8, 10},
				Interval{15, 18},
			}},
			ans{[]Interval{
				Interval{1, 6},
				Interval{8, 10},
				Interval{15, 18},
			}},
		},

		question{
			para{[]Interval{
				Interval{1, 4},
				Interval{4, 5},
			}},
			ans{[]Interval{
				Interval{1, 5},
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(ans.one, merge(para.one), "输入 %v", para.one)
	}
}
