package problem387

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	s string
	ans int
}{

	{
		"",
		-1,
	},

	{
		"oo",
		-1,
	},

	{
		"leetcode",
		0,
	},

	{
		"loveleetcode",
		2,
	},
}

func Test_firstUniqChar(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, firstUniqChar(q.s), "输入 %s", q.s)
	}
}

func Benchmark_firstUniqChar(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			firstUniqChar(q.s)
		}		
	}
}