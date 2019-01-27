package problem212

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var questions = []struct {
	board [][]byte
	words []string
	ans   []string
}{
	{
		[][]byte{
			[]byte{'a'},
			[]byte{'a'},
		},
		[]string{"a"},
		[]string{"a"},
	},

	{
		[][]byte{
			[]byte{'a'},
		},
		[]string{"a", "a"},
		[]string{"a"},
	},

	{
		[][]byte{
			[]byte{'o', 'a', 'a', 'n'},
			[]byte{'e', 't', 'a', 'e'},
			[]byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'i', 'v'},
		},
		[]string{"oath", "pea", "eat", "rain"},
		[]string{"oath", "eat"},
	},
}

func Test_findWords(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(findWords(q.board, q.words), q.ans)
	}
}

func Test_findWords2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range questions {
		ast.Equal(findWords2(q.board, q.words), q.ans, "输入 %v", q.words)
	}
}

func Benchmark_findWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findWords(q.board, q.words)
		}
	}
}

func Benchmark_findWords2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, q := range questions {
			findWords2(q.board, q.words)
		}
	}
}
