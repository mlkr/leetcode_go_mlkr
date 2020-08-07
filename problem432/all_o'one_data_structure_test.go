package problem432

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AllOne(t *testing.T) {
	ast := assert.New(t)

	ao := Constructor()
	ast.Equal("", ao.GetMaxKey())
	ast.Equal("", ao.GetMinKey())

	ao.Inc("m")
	ao.Inc("l")
	ao.Inc("m")
	ao.Inc("m")
	ao.Inc("l")

	ast.Equal("m", ao.GetMaxKey())
	ast.Equal("l", ao.GetMinKey())

	ao.Dec("m")
	ao.Dec("m")
	ao.Dec("a")

	ast.Equal("l", ao.GetMaxKey())
	ast.Equal("m", ao.GetMinKey())

	ao.Dec("m")

	ast.Equal("l", ao.GetMaxKey())
	ast.Equal("l", ao.GetMinKey())

	bo := Constructor()
	bo.Inc("m")
	bo.Inc("l")
	bo.Inc("m")
	bo.Inc("m")
	bo.Inc("l")
	bo.Inc("a")

	ast.Equal("m", bo.GetMaxKey())
	ast.Equal("a", bo.GetMinKey())

	co := Constructor()
	co.Inc("m")
	co.Inc("m")
	co.Inc("m")
	co.Inc("l")
	co.Inc("l")

	ast.Equal("m", co.GetMaxKey())
	ast.Equal("l", co.GetMinKey())

	do := Constructor()
	do.Inc("a")
	do.Inc("b")
	do.Inc("b")
	do.Inc("c")
	do.Inc("c")
	do.Inc("d")
	do.Inc("d")
	do.Inc("d")
	do.Inc("d")

	ast.Equal("d", do.GetMaxKey())
	ast.Equal("a", do.GetMinKey())

	eo := Constructor()
	eo.Inc("a")
	eo.Inc("a")
	eo.Dec("a")
	eo.Dec("a")

	ast.Equal("", eo.GetMaxKey())
	ast.Equal("", eo.GetMinKey())

}

func Benchmark_AllOne(b *testing.B) {
	ao := Constructor()

	for i := 0; i < b.N; i++ {
		ao.Inc("m")
		ao.Inc("m")
		ao.Inc("m")
		ao.Inc("l")
		ao.Inc("l")

		ao.GetMaxKey()
		ao.GetMinKey()

		ao.Dec("m")
		ao.Dec("m")
		ao.Dec("a")

		ao.GetMaxKey()
		ao.GetMinKey()

	}
}
