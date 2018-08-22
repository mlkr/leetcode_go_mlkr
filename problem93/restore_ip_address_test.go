package problem93

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one []string
}

type question struct {
	para
	ans
}

func TestRestoreIpAddresses(t *testing.T) {
	ast := assert.New(t)

	questions := []question{

		question{
			para{
				"172162541",
			},
			ans{[]string{
				"17.216.25.41",
				"17.216.254.1",
				"172.16.25.41",
				"172.16.254.1",
				"172.162.5.41",
				"172.162.54.1",
			}},
		},

		question{
			para{
				"25525511135",
			},
			ans{[]string{
				"255.255.11.135",
				"255.255.111.35",
			}},
		},
	}

	for _, q := range questions {
		para, ans := q.para, q.ans
		ast.Equal(restoreIpAddresses(para.one), ans.one, "输入 %v", para)
	}
}
