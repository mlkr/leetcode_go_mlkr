package problem336

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	words []string
	ans   [][]int
}{
	{
		[]string{"a"},
		[][]int{},
	},

	{
		[]string{"abcd", "dcba", "lls", "s", "sssll"},
		[][]int{
			[]int{0, 1},
			[]int{1, 0},
			[]int{3, 2},
			[]int{2, 4},
		},
	},

	{
		[]string{"bat", "tab", "cat"},
		[][]int{
			[]int{0, 1},
			[]int{1, 0},
		},
	},
}

func Test_palindromePairs(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.ElementsMatch(q.ans, palindromePairs(q.words))
	}
}

func Benchmark_palindromePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			palindromePairs(q.words)
		}
	}
}
