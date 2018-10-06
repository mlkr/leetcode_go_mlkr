package problem127

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
		step      int
	}{
		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			5,

			// 1				hit
			// 2			hot
			// 3		dot		lot
			// 4	dog
			// 5 cog	log
		},

		{
			"hot",
			"dog",
			[]string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"},
			3,
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog", "dot"},
			3,
		},

		{
			"talk",
			"tail",
			[]string{"talk", "tons", "fall", "tail", "gale", "hall", "negs"},
			0,
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog"},
			0,
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log"},
			0,
		},
	}

	for _, q := range questions {
		ast.Equal(ladderLength(q.beginWord, q.endWord, q.wordList), q.step, "输入 %v", q)
		ast.Equal(ladderLength2(q.beginWord, q.endWord, q.wordList), q.step, "输入 %v", q)
		ast.Equal(ladderLength3(q.beginWord, q.endWord, q.wordList), q.step, "输入 %v", q)
	}
}
