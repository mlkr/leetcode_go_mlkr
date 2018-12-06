package problem187

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	para string
	ans  []string
}{
	{
		"AAAAAAAAAAAA",
		[]string{
			"AAAAAAAAAA",
		},
	},

	{
		"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
		[]string{
			"AAAAACCCCC",
			"CCCCCAAAAA",
		},
	},
}

func Test_findRepeatedDnaSequences(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(findRepeatedDnaSequences(q.para), q.ans)
	}
}

func Benchmark_findRepeatedDnaSequences(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findRepeatedDnaSequences(q.para)
		}
	}
}
