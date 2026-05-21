package problem523

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	strs []string
	ans  int
}{
	{
		[]string{"aba", "cdc", "eae"},
		3,
	},

	{
		[]string{"aaa", "aaa", "aa"},
		-1,
	},

	{
		[]string{"a", "aaa", "aaa", "aa", "aa"},
		-1,
	},

	{
		[]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"},
		6,
	},
}

func Test_findLUSlength(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		res := findLUSlength(q.strs)
		ast.Equal(q.ans, res, "输入 %v", q.strs)
	}
}

func Benchmark_findLUSlength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findLUSlength(q.strs)
		}
	}
}
