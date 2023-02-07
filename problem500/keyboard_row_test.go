package problem500

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	words, ans []string
}{
	{
		[]string{"Hello", "Alaska", "Dad", "Peace"},
		[]string{"Alaska", "Dad"},
	},

	{
		[]string{"omk"},
		[]string{},
	},

	{
		[]string{"adsdf", "sfd"},
		[]string{"adsdf", "sfd"},
	},
}

func Test_findDiagonalOrder(t *testing.T) {
	ast := assert.New(t)
	for _, q := range questions {
		ast.Equal(q.ans, findWords(q.words), "输入 %v", q.words)
	}
}

func Benchmark_findWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findWords(q.words)
		}
	}
}
