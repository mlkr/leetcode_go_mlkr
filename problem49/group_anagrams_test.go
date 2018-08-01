package problem49

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []string
}

type ans struct {
	one [][]string
}

type question struct {
	para
	ans
}

func Test_groupAnagrams(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		question{
			para{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans{[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			}},
		},
	}

	for _, q := range questions {
		ans, para := q.ans, q.para
		res := groupAnagrams(para.one)

		for _, v := range res {
			ast.Equal(ans.one[len(v)-1], v, "输入 %v", v)
		}
	}
}
