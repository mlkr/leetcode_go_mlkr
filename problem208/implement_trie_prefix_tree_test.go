package problem208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	ast := assert.New(t)

	trie := Constructor()
	trie.Insert("apple")
	ast.True(trie.Search("apple"))
	ast.False(trie.Search("app"))
	ast.True(trie.StartsWith("app"))

	trie.Insert("app")
	ast.True(trie.Search("app"))
}

func Benchmark_Trie(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trie := Constructor()
		trie.Insert("apple")
		trie.Search("apple")
		trie.Search("app")
		trie.StartsWith("app")

		trie.Insert("app")
		trie.Search("app")
	}
}
