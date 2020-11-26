package problem460

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LFUCache(t *testing.T) {
	ast := assert.New(t)

	lfu := Constructor(2)

	lfu.Put(1, 1)
	lfu.Put(2, 2)
	ast.Equal(1, lfu.Get(1), "输入 1")

	lfu.Put(3, 3)
	ast.Equal(-1, lfu.Get(2), "输入 2")
	ast.Equal(3, lfu.Get(3), "输入 3")

	lfu.Put(4, 4)
	ast.Equal(3, lfu.Get(3), "输入 3")
	ast.Equal(4, lfu.Get(4), "输入 4")

	lfu.Put(4, 5)
	ast.Equal(5, lfu.Get(4), "输入 5")

	lfu2 := Constructor(0)
	lfu2.Put(0, 0)
	ast.Equal(-1, lfu2.Get(0))

	lfu3 := Constructor(2)
	lfu3.Put(2, 1)
	lfu3.Put(3, 2)
	ast.Equal(2, lfu3.Get(3))
	ast.Equal(1, lfu3.Get(2))

	lfu3.Put(4, 3)
	ast.Equal(1, lfu3.Get(2))
	ast.Equal(-1, lfu3.Get(3))
	ast.Equal(3, lfu3.Get(4))

}

func Benchmark_LFUCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lfu := Constructor(2)

		lfu.Put(1, 1)
		lfu.Put(2, 2)
		lfu.Get(1)

		lfu.Put(3, 3)
		lfu.Get(2)
		lfu.Get(3)

		lfu.Put(4, 4)
		lfu.Get(3)
		lfu.Get(4)
	}
}
