package problem146

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lruCache(t *testing.T) {
	ast := assert.New(t)

	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	ast.Equal(cache.Get(1), 1)

	cache.Put(3, 3)
	ast.Equal(cache.Get(2), -1)

	cache.Put(4, 4)
	ast.Equal(cache.Get(1), -1)
	ast.Equal(cache.Get(3), 3)
	ast.Equal(cache.Get(4), 4)
}

func Benchmark_lruCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Get(1)

		cache.Put(3, 3)
		cache.Get(2)

		cache.Put(4, 4)
		cache.Get(1)
		cache.Get(3)
		cache.Get(4)
	}
}
