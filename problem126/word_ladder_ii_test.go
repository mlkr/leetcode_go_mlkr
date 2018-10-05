package problem126

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	ast := assert.New(t)

	questions := []struct {
		beginWord string
		endWord   string
		wordList  []string
		step      [][]string
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{
				[]string{"hit", "hot", "dot", "dog", "cog"},
				[]string{"hit", "hot", "lot", "log", "cog"},
			},

			// 1				hit
			// 2			hot
			// 3		dot		lot
			// 4	dog       log
			// 5 cog       cog
		},

		{
			"a",
			"c",
			[]string{"a", "b", "c"},
			[][]string{
				[]string{"a", "c"},
			},
		},

		{
			"talk",
			"tail",
			[]string{"talk", "tons", "fall", "tail", "gale", "hall", "negs"},
			[][]string{},
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog"},
			[][]string{},
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			[][]string{},
		},
	}

	for _, q := range questions {
		ast.Equal(findLadders(q.beginWord, q.endWord, q.wordList), q.step, "输入 %v", q)
	}
}
