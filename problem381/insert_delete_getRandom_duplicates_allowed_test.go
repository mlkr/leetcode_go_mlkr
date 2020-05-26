package problem381

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomizedCollection(t *testing.T) {
	ast := assert.New(t)

	rc := Constructor()
	ast.True(rc.Insert(1))
	ast.False(rc.Insert(1))
	ast.True(rc.Insert(2))
	rc.GetRandom()
	ast.True(rc.Remove(1))
	rc.GetRandom()
	ast.True(rc.Remove(2))
	ast.False(rc.Remove(2))
}

func benchmark_RandomizedCollection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rc := Constructor()
		rc.Insert(1)
		rc.Insert(1)
		rc.GetRandom()
		rc.Remove(1)
		rc.GetRandom()

	}
}
