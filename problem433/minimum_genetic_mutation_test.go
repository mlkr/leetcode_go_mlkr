package problem433

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct{
	start,end string
	bank []string
	ans int
}{
	{
		"AACCGGTT",
		"AACCGGTA",
		[]string{"AACCGGTA"},
		1,
	},

	{
		"AACCGGTT",
		"AAACGGTA",
		[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
		2,
	},

	{
		"AAAAACCC",
		"AACCCCCC",
		[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
		3,
	},

	{
		"AAAAACCC",
		"ACCCCCC",
		[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
		-1,
	},
}

func Test_minMutation(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(q.ans, minMutation(q.start,q.end,q.bank), "start %v", q.start)
	}
}

func Benchmark_minMutation(b *testing.B) {
	for i:=0;i<b.N;i++{
		for _, q := range questions {
			minMutation(q.start,q.end,q.bank)
		}
	}
}