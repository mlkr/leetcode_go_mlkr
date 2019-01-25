package problem211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordDictionary(t *testing.T) {
	ast := assert.New(t)

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	ast.False(wd.Search("pad"))
	ast.True(wd.Search("bad"))
	ast.True(wd.Search(".ad"))
	ast.True(wd.Search("b.."))

	wd2 := Constructor()
	wd2.AddWord("a")
	wd2.AddWord("ab")
	ast.True(wd2.Search("a"), "a")
	ast.True(wd2.Search("a."), "a.")
	ast.True(wd2.Search("ab"), "ab")
	ast.False(wd2.Search(".a"), ".a")
	ast.True(wd2.Search(".b"), ".b")
	ast.False(wd2.Search("ab."), "ab.")
	ast.True(wd2.Search("."), ".")
	ast.True(wd2.Search(".."), "..")

	wd3 := Constructor()
	wd3.AddWord("at")
	wd3.AddWord("and")
	wd3.AddWord("an")
	wd3.AddWord("add")
	ast.False(wd3.Search("a"), "a")
	ast.False(wd3.Search(".at"), ".at")
	wd3.AddWord("bat")
	ast.True(wd3.Search(".at"), ".at")
	ast.True(wd3.Search("an."), "an.")
	ast.False(wd3.Search("a.d."), "a.d.")
	ast.False(wd3.Search("b."), "b.")
	ast.True(wd3.Search("a.d"), "a.d")
	ast.False(wd3.Search("."), ".")

}

func Benchmark_WordDictionary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wd := Constructor()
		wd.AddWord("bad")
		wd.AddWord("dad")
		wd.AddWord("mad")
		wd.Search("pad")
		wd.Search("bad")
		wd.Search(".ad")
		wd.Search("b..")

	}
}
