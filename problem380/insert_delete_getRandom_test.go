package problem380

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomizedSet(t *testing.T) {
	ast := assert.New(t)

	rs := Constructor()
	ast.True(rs.Insert(1))
	ast.False(rs.Remove(2))
	ast.True(rs.Insert(2))

	n := rs.GetRandom()
	fmt.Println(n)

	ast.True(rs.Remove(1))
	ast.False(rs.Insert(2))
	ast.Equal(2, rs.GetRandom())
}

func Benchmark_RandomizedSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rs := Constructor()
		rs.Insert(1)
		rs.Remove(2)
		rs.Insert(2)

		rs.GetRandom()

		rs.Remove(1)
		rs.Insert(2)
		rs.GetRandom()
	}
}
